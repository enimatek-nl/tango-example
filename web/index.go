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

type IndexController struct {
	Todos     []server.Todo `tng:"todos"`
	Busy      bool          `tng:"busy"`
	Modal     bool          `tng:"modal"`
	ModalInfo string        `tng:"info"`
	Add       tango.SFunc   `tng:"add"`
	Edit      tango.SFunc   `tng:"edit"`
	Delete    tango.SFunc   `tng:"delete"`
}

func (i IndexController) Config() tango.ComponentConfig {
	return tango.ComponentConfig{
		Name:   "IndexController",
		Kind:   tango.Controller,
		Scoped: false,
	}
}

func (i *IndexController) Constructor(tng tango.Hook) bool {
	i.Modal = false
	i.ModalInfo = ""
	i.Todos = []server.Todo{}
	i.Busy = true

	i.Add = func(loc *tango.Hook) {
		loc.Self.Nav("/edit/0")
	}

	i.Edit = func(loc *tango.Hook) {
		if id, found := loc.Get("todo.ID"); found {
			loc.Self.Nav("/edit/" + fmt.Sprintf("%d", id.Int()))
		}
	}

	i.Delete = func(loc *tango.Hook) {
		i.Busy = true
		go func() {
			if id, found := loc.Get("todo.ID"); found {
				req, _ := http.NewRequest(http.MethodDelete, "/api/todo?id="+fmt.Sprintf("%d", id.Int()), nil)
				if _, err := http.DefaultClient.Do(req); err == nil {
					refresh(i, tng.Scope)
				} else {
					i.Modal = true
					i.ModalInfo = err.Error()
				}
				i.Busy = false
				tng.Absorb(i)
			} else {
				println("todo ID not found")
			}
		}()
		tng.Absorb(i)
	}

	tng.Absorb(i)
	return true
}

func (i IndexController) BeforeRender(tng tango.Hook) {}

func (i *IndexController) AfterRender(tng tango.Hook) {
	go refresh(i, tng.Scope)
}

func refresh(i *IndexController, s *tango.Scope) {
	if resp, err := http.Get("/api/todo"); err == nil {
		defer resp.Body.Close()

		// load all todos from api-backend
		json.NewDecoder(resp.Body).Decode(&i.Todos)

		i.Busy = false

		s.Absorb(i)
	}
}

func (i IndexController) Render() string {
	return tmplIndex
}
