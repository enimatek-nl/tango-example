package web

import (
	"github.com/enimatek-nl/tango"
	"syscall/js"
)

const SHOW = "show"

type Busy struct{}

func (l Busy) Config() tango.ComponentConfig {
	return tango.ComponentConfig{
		Name:   "Busy",
		Kind:   tango.Tag,
		Scoped: false,
	}
}

func (l Busy) Hook(self *tango.Tango, scope *tango.Scope, hook tango.ComponentHook, attrs map[string]string, node js.Value, queue *tango.Queue) bool {
	if v, e := attrs[SHOW]; e {
		scope.Subscribe(
			v,
			func(scope *tango.Scope, value js.Value) {
				if value.Bool() {
					node.Get("style").Set("display", "block")
				} else {
					node.Get("style").Set("display", "none")
				}
			},
		)
	} else {
		panic("don't forget to set the '" + SHOW + "' attr")
	}
	node.Get("style").Set("display", "none")
	return true
}

func (l Busy) Render() string {
	return `
            <div class="loading">
                <img width="64" height="64" src="gifs/loading.gif"></img>
            </div>
`
}