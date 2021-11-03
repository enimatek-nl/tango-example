package web

import (
	_ "embed"
	"github.com/enimatek-nl/tango"
	"syscall/js"
)

//go:embed edit.html
var tmplEdit string

type EditController struct{}

func (i EditController) Config() tango.ComponentConfig {
	return tango.ComponentConfig{
		Name:   "EditController",
		Kind:   tango.Controller,
		Scoped: false,
	}
}

func (i *EditController) cancel(value js.Value, scope *tango.Scope) {
	i.busy(scope, true)
}

func (i *EditController) save(value js.Value, scope *tango.Scope) {
	i.busy(scope, true)
}

func (i *EditController) busy(scope *tango.Scope, bsy bool) {
	scope.Set("busy", js.ValueOf(bsy))
	scope.Digest()
}

func (i EditController) Hook(self *tango.Tango, scope *tango.Scope, hook tango.ComponentHook, attrs map[string]string, node js.Value, queue *tango.Queue) bool {
	switch hook {
	case tango.Construct:
		scope.SetFunc("save", i.save)
		scope.SetFunc("cancel", i.cancel)
	}
	return true
}

func (i EditController) Render() string {
	return tmplEdit
}
