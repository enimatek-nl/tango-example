package web

import (
	"github.com/enimatek-nl/tango"
	"syscall/js"
)

const CONTENT = "content"

type Modal struct{}

func (m Modal) Config() tango.ComponentConfig {
	return tango.ComponentConfig{
		Name:   "Modal",
		Kind:   tango.Tag,
		Scoped: false,
	}
}

func (m Modal) Constructor(hook tango.Hook) bool {
	if v, e := hook.Attrs[CONTENT]; e {
		hook.Scope.Subscribe(v, func(scope *tango.Scope, value js.Value) {
			hook.Scope.Set("content", value)
		})
	} else {
		panic("don't forget to set the '" + CONTENT + "' attr")
	}

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

	js.Global().Get("window").Set("onclick", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		hook.Node.Get("style").Set("display", "none")
		return nil
	}))

	return true
}

func (m Modal) BeforeRender(hook tango.Hook) {}

func (m Modal) AfterRender(hook tango.Hook) {}

func (m Modal) Render() string {
	return `
<div id="myModal" class="modal">
  <div class="modal-content">
    <p tng-bind="content"></p>
  </div>
</div>`
}
