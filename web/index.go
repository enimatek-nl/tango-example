package web

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/enimatek-nl/tango"
	"github.com/enimatek-nl/tango-example/server"
	"net/http"
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

func (i *IndexController) Constructor(tng tango.Hook) bool {
	// functions
	tng.Set("todos", []server.Todo{})
	tng.Set("busy", false)
	tng.SetFunc("add", func(loc *tango.Hook) {
		loc.Self.Nav("/edit/0")
	})
	tng.SetFunc("edit", func(loc *tango.Hook) {
		if id, found := loc.Get("todo.ID"); found {
			loc.Self.Nav("/edit/" + fmt.Sprintf("%d", id.Int()))
		}
	})
	tng.SetFunc("delete", func(loc *tango.Hook) {
		tng.Set("busy", true)
		go func() {
			if id, found := loc.Get("todo.ID"); found {
				req, _ := http.NewRequest(http.MethodDelete, "/api/todo?id="+fmt.Sprintf("%d", id.Int()), nil)
				if _, err := http.DefaultClient.Do(req); err == nil {
					refresh(&tng)
				} else {
					tng.Set("info", err.Error())
					tng.Set("modal", true)
				}
				tng.Set("busy", false)
				tng.Scope.Digest()
			} else {
				println("todo ID not found")
			}
		}()
	})
	return true
}

func (i IndexController) BeforeRender(tng tango.Hook) {}

func (i IndexController) AfterRender(tng tango.Hook) {
	tng.Set("busy", true)
	go refresh(&tng)
}

func refresh(tng *tango.Hook) {
	if resp, err := http.Get("/api/todo"); err == nil {
		defer resp.Body.Close()

		// load all todos from api-backend
		var todos []server.Todo
		json.NewDecoder(resp.Body).Decode(&todos)
		tng.Set("todos", todos)

		// remove loading overlay
		tng.Set("busy", false)
		tng.Scope.Digest()
	}
}

func (i IndexController) Render() string {
	return tmplIndex
}
