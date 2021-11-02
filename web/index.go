package web

import (
	_ "embed"
	"github.com/enimatek-nl/tango"
	"syscall/js"
)

//go:embed index.html
var tmplIndex string

type IndexController struct {}

func (i IndexController) Config() tango.ComponentConfig {
	return tango.ComponentConfig{
		Name:   "IndexController",
		Kind:   tango.Controller,
		Scoped: false,
	}
}

func (i *IndexController) newTodo(value js.Value, scope *tango.Scope) {
	scope.Set("busy", js.ValueOf(true))
}

func (i IndexController) Hook(self *tango.Tango, scope *tango.Scope, hook tango.ComponentHook, attrs map[string]string, node js.Value, queue *tango.Queue) bool {
	switch hook {
	case tango.Construct:
		scope.Set("busy", js.ValueOf(false))
		scope.SetFunc("newTodo", i.newTodo)
	}
	return true
}

func (i IndexController) Render() string {
	return tmplIndex
}
