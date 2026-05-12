package template

import "testing"

func TestTemplateModule(t *testing.T) {
	modInterface := GetInstance()
	mod, ok := modInterface.(*TemplateModule)
	if !ok {
		t.Fatal("cannot cast IModule to TemplateModule")
	}

	if err := mod.Init(); err != nil {
		t.Error(err)
	}
}
