package template

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var groupPath = "/templates"

type TestHandlerConfig struct {
	name               string
	whenPath           string
	expectDescription  string
	expectStatus       int
	method             string
	expectResponseBody any
	inputBody          any
}

type TestResponse[T any] struct {
	Status      int    `json:"status"`
	Description string `json:"message"`
	Data        T      `json:"data,omitempty"`
}

type testServeMux struct {
	handlers map[string]map[string]http.HandlerFunc
}

func newTestServeMux() *testServeMux {
	return &testServeMux{handlers: make(map[string]map[string]http.HandlerFunc)}
}

func (m *testServeMux) handle(method, path string, handler http.HandlerFunc) {
	if m.handlers[path] == nil {
		m.handlers[path] = make(map[string]http.HandlerFunc)
	}
	m.handlers[path][method] = handler
}

func (m *testServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handlers, ok := m.handlers[r.URL.Path]; ok {
		if h, ok := handlers[r.Method]; ok {
			h(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func TestHandler(t *testing.T) {
	testConfigFile(t)
	mod := testCreateModule(t)

	mux := newTestServeMux()
	for _, r := range *mod.GetRoutes() {
		mux.handle(r.Method, groupPath+r.Path, r.Handler)
	}

	testPathHandler(t, mux, mod)
}

func jsonMarshalData(input any) (io.Reader, error) {
	b, err := json.Marshal(&input)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func testPathHandler(t *testing.T, mux *testServeMux, mod *TemplateModule) {
	mockCode := Template{Code: "test_template"}
	active := true
	mockCreate := Template{Code: mockCode.Code, Name: "Test Template", Active: &active}
	mockUpdate := Template{Code: mockCode.Code, Name: "Test Template updated", Active: &active}
	mockEmpty := ""

	testCases := []TestHandlerConfig{
		{
			name:              "Normal list",
			method:            http.MethodGet,
			whenPath:          "/search?keyword=hello&page=1&limit=10",
			expectDescription: "Success",
			expectStatus:      http.StatusOK,
			inputBody:         mockEmpty,
		},
		{
			name:              "Create template",
			method:            http.MethodPost,
			whenPath:          "",
			expectDescription: "Success",
			expectStatus:      http.StatusOK,
			inputBody:         mockCreate,
		},
		{
			name:               "Get template by code",
			method:             http.MethodPost,
			whenPath:           "/info",
			expectDescription:  "Success",
			expectStatus:       http.StatusOK,
			inputBody:          mockCode,
			expectResponseBody: mockCreate,
		},
		{
			name:              "Create duplicate template",
			method:            http.MethodPost,
			whenPath:          "",
			expectDescription: fmt.Sprintf("Template code %s already exist.", mockCreate.Code),
			expectStatus:      http.StatusBadRequest,
			inputBody:         mockCreate,
		},
		{
			name:              "Delete template",
			method:            http.MethodDelete,
			whenPath:          "",
			expectDescription: "Success",
			expectStatus:      http.StatusOK,
			inputBody:         mockCreate,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reqPath := fmt.Sprintf("%s%s", groupPath, tc.whenPath)
			bodyIo, err := jsonMarshalData(tc.inputBody)
			if err != nil {
				assert.NoError(t, err)
			}
			req := httptest.NewRequest(tc.method, reqPath, bodyIo)
			rec := httptest.NewRecorder()
			mux.ServeHTTP(rec, req)

			resp := TestResponse[Template]{}
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
			assert.Equal(t, tc.expectStatus, rec.Code)
			assert.Equal(t, tc.expectDescription, resp.Description)
		})
	}
}
