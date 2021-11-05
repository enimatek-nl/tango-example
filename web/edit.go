package web

import (
	_ "embed"
	"github.com/enimatek-nl/tango"
	"syscall/js"
)

//go:embed edit.html
var tmplEdit string

type EditController struct{}

func (e EditController) Config() tango.ComponentConfig {
	return tango.ComponentConfig{
		Name:   "EditController",
		Kind:   tango.Controller,
		Scoped: false,
	}
}

func (e EditController) Constructor(hook tango.Hook) bool {
	hook.Scope.SetFunc("cancel", func(value js.Value, scope *tango.Scope) {
		hook.Self.Nav("/")
	})
	return true
}

func (e EditController) BeforeRender(hook tango.Hook) {}

func (e EditController) AfterRender(hook tango.Hook) {}

func (e EditController) Render() string {
	return tmplEdit
}
