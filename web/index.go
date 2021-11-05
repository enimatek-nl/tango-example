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

func (i IndexController) Constructor(hook tango.Hook) bool {
	// variables
	hook.Scope.Set("busy", true) // enable loading overlay
	hook.Scope.Set("todos", []server.Todo{})
	// functions
	hook.Scope.SetFunc("add", func(value js.Value, scope *tango.Scope) {
		hook.Self.Nav("/edit/0")
	})
	hook.Scope.SetFunc("edit", func(value js.Value, scope *tango.Scope) {
		if id, found := scope.Get("todo.ID"); found {
			hook.Self.Nav("/edit/" + id.String())
		}
	})
	return true
}

func (i IndexController) BeforeRender(hook tango.Hook) {}

func (i IndexController) AfterRender(hook tango.Hook) {
	go func() {
		if resp, err := http.Get("/api/get"); err == nil {
			var todos []server.Todo
			// load all todos from api-backend
			defer resp.Body.Close()
			json.NewDecoder(resp.Body).Decode(&todos)
			hook.Scope.Set("todos", todos)
			// remove loading overlay
			hook.Scope.Set("busy", false)
			hook.Scope.Digest()
		}
	}()
}

func (i IndexController) Render() string {
	return tmplIndex
}
