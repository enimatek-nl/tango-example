package web

import (
	"github.com/enimatek-nl/tango"
	"syscall/js"
)

const SHOW = "show"

type Busy struct{}

func (b Busy) Config() tango.ComponentConfig {
	return tango.ComponentConfig{
		Name:   "Busy",
		Kind:   tango.Tag,
		Scoped: false,
	}
}

func (b Busy) Constructor(hook tango.Hook) bool {
	if v, e := hook.Attrs[SHOW]; e {
		hook.Scope.Subscribe(v, func(scope *tango.Scope, value js.Value) {
			if value.Bool() {
				hook.Node.Get("style").Set("display", "block")
			} else {
				hook.Node.Get("style").Set("display", "none")
			}
		})
	} else {
		panic("don't forget to set the '" + SHOW + "' attr")
	}
	hook.Node.Get("style").Set("display", "none")
	return true
}

func (b Busy) BeforeRender(hook tango.Hook) {}

func (b Busy) AfterRender(hook tango.Hook) {}

func (b Busy) Render() string {
	return `
            <div class="loading">
                <img width="64" height="64" src="loading.gif"></img>
            </div>
`
}
