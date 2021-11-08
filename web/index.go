package web

import (
	_ "embed"
	"encoding/json"
	"fmt"
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
			hook.Self.Nav("/edit/" + fmt.Sprintf("%d", id.Int()))
		}
	})
	hook.Scope.SetFunc("delete", func(value js.Value, scope *tango.Scope) {
		if id, found := scope.Get("todo.ID"); found {
			hook.Scope.Set("busy", true)
			go func() {
				req, _ := http.NewRequest(http.MethodDelete, "/api/todo?id="+fmt.Sprintf("%d", id.Int()), nil)
				if _, err := http.DefaultClient.Do(req); err == nil {
					refresh(&hook)
				} else {
					println(err)
				}
			}()
		}
	})
	return true
}

func (i IndexController) BeforeRender(hook tango.Hook) {}

func (i IndexController) AfterRender(hook tango.Hook) {
	go refresh(&hook)
}

func refresh(hook *tango.Hook) {
	if resp, err := http.Get("/api/todo"); err == nil {
		defer resp.Body.Close()

		// load all todos from api-backend
		var todos []server.Todo
		json.NewDecoder(resp.Body).Decode(&todos)
		hook.Scope.Set("todos", todos)

		// remove loading overlay
		hook.Scope.Set("busy", false)
		hook.Scope.Digest()
	}
}

func (i IndexController) Render() string {
	return tmplIndex
}
