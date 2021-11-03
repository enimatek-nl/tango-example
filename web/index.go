package web

import (
	_ "embed"
	"encoding/json"
	"github.com/enimatek-nl/tango"
	"github.com/enimatek-nl/tango-example/server"
	"net/http"
	"syscall/js"
)

//go:embed index.html
var tmplIndex string

type IndexController struct{}

func (i IndexController) Config() tango.ComponentConfig {
	return tango.ComponentConfig{
		Name:   "IndexController",
		Kind:   tango.Controller,
		Scoped: false,
	}
}

func (i IndexController) Hook(self *tango.Tango, scope *tango.Scope, hook tango.ComponentHook, attrs map[string]string, node js.Value, queue *tango.Queue) bool {
	switch hook {
	case tango.Construct:
		scope.SetFunc("add", func(value js.Value, scope *tango.Scope) {
			scope.Nav("/edit/0")
		})
		scope.Set("busy", js.ValueOf(true))
		scope.Set("todos", js.ValueOf([]interface{}{}))
	case tango.AfterRender:
		go func() {
			if resp, err := http.Get("/api/get"); err == nil {
				var todos []server.Todo
				// load all todos from api-backend
				defer resp.Body.Close()
				json.NewDecoder(resp.Body).Decode(&todos)

				println(todos[0].Title)

				y := make([]interface{}, len(todos))
				for i, v := range todos {
					y[i] = js.ValueOf(v)
				}
				scope.Set("todos", js.ValueOf(y))

				// remove loading screen
				scope.Set("busy", js.ValueOf(false))
				scope.Digest()
			}
		}()
	}
	return true
}

func (i IndexController) Render() string {
	return tmplIndex
}
