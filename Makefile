BINARY := app
SERVICE_NAME := my-service

.PHONY: all build run dev local air tidy test clean new swagger

all: build

build:
	@echo "-- building"
	go build -o $(BINARY) .

run:
	./$(BINARY)

build-run: build run

dev:
	go run main.go --env=dev --loadEnvFile=true --loadConfFile=false

local:
	go run main.go --env=local --loadEnvFile=true --loadConfFile=true

air:
	air

test:
	@echo "-- testing"
	go test ./... -v -cover

swagger:
	@echo "-- generating swagger docs"
	swag init --generalInfo main.go --output docs

tidy:
	go mod tidy

clean:
	@echo "-- cleaning"
	-rm -f $(BINARY) app.db
	-rm -rf tmp/

new:
	@echo "-- Creating new module: $(MODULE_NAME)"
	@if [ -z "$(MODULE_NAME)" ]; then echo "Usage: make new MODULE_NAME=foo CAP_MODULE_NAME=Foo"; exit 1; fi
	cp -r ./.static/templates/module-template ./.static/templates/$(MODULE_NAME)
	find ./.static/templates/$(MODULE_NAME) -type f -name "*.go" -exec sed -i '' 's/template/$(MODULE_NAME)/g' {} +
	find ./.static/templates/$(MODULE_NAME) -type f -name "*.go" -exec sed -i '' 's/Template/$(CAP_MODULE_NAME)/g' {} +
	mv ./.static/templates/$(MODULE_NAME) ./module/$(MODULE_NAME)
	@echo "-- Created module at ./module/$(MODULE_NAME)"
