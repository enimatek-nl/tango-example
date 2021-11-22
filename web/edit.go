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

type EditController struct {
	Busy   bool        `tng:"busy"`
	Todo   server.Todo `tng:"todo"`
	Cancel tango.SFunc `tng:"cancel"`
	Save   tango.SFunc `tng:"save"`
}

func (e EditController) Config() tango.ComponentConfig {
	return tango.ComponentConfig{
		Name:   "EditController",
		Kind:   tango.Controller,
		Scoped: false,
	}
}

func (e *EditController) Constructor(hook tango.Hook) bool {
	e.Busy = true
	e.Todo = server.Todo{}

	e.Cancel = func(hook *tango.Hook) {
		hook.Self.Nav("/")
	}

	e.Save = func(hook *tango.Hook) {
		hook.Scope.Set("busy", true)
		go func() {
			r := strings.NewReader(
				hook.Scope.GetJSON("todo"),
			)
			http.Post("/api/todo", "application/json", r)
			hook.Self.Nav("/")
		}()
	}

	hook.Absorb(e)
	return true
}

func (e EditController) BeforeRender(hook tango.Hook) {}

func (e *EditController) AfterRender(hook tango.Hook) {
	go func() {
		if i, o := hook.Attrs["id"]; o {
			if resp, err := http.Get("/api/todo?id=" + i); err == nil {
				defer resp.Body.Close()
				json.NewDecoder(resp.Body).Decode(&e.Todo)
			}
		}
		e.Busy = false
		hook.Absorb(e)
	}()
}

func (e EditController) Render() string {
	return tmplEdit
}
