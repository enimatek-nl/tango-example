package web

import (
	_ "embed"
	"encoding/json"
	"github.com/enimatek-nl/tango"
	"github.com/enimatek-nl/tango-example/server"
	"net/http"
	"strings"
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
	hook.Scope.Set("busy", false)
	hook.Scope.Set("todo", server.Todo{})
	hook.Scope.SetFunc("cancel", func(hook *tango.Hook) {
		hook.Self.Nav("/")
	})
	hook.Scope.SetFunc("save", func(hook *tango.Hook) {
		hook.Scope.Set("busy", true)
		go func() {
			r := strings.NewReader(
				hook.Scope.GetJSON("todo"),
			)
			http.Post("/api/todo", "application/json", r)
			hook.Self.Nav("/")
		}()
	})
	return true
}

func (e EditController) BeforeRender(hook tango.Hook) {}

func (e EditController) AfterRender(hook tango.Hook) {
	go func() {
		if i, e := hook.Attrs["id"]; e {
			if resp, err := http.Get("/api/todo?id=" + i); err == nil {
				var todo server.Todo
				defer resp.Body.Close()
				json.NewDecoder(resp.Body).Decode(&todo)
				hook.Scope.Set("todo", todo)
				hook.Scope.Digest()
			}
		}
	}()
}

func (e EditController) Render() string {
	return tmplEdit
}
