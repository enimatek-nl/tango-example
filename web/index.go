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

	i.Add = func(self *tango.Tango, this js.Value, local *tango.Scope) {
		self.Nav("/edit/0")
	}

	i.Edit = func(self *tango.Tango, this js.Value, local *tango.Scope) {
		if id, found := local.Get("todo.ID"); found {
			self.Nav("/edit/" + fmt.Sprintf("%d", id.Int()))
		}
	}

	i.Delete = func(self *tango.Tango, this js.Value, local *tango.Scope) {
		i.Busy = true
		go func() {
			if id, found := local.Get("todo.ID"); found {
				req, _ := http.NewRequest(http.MethodDelete, "/api/todo?id="+fmt.Sprintf("%d", id.Int()), nil)
				if _, err := http.DefaultClient.Do(req); err == nil {
					refresh(i, tng)
				} else {
					i.Modal = true
					i.ModalInfo = err.Error()
				}
				i.Busy = false
				tng.Digest(i)
			} else {
				println("todo ID not found")
			}
		}()
		tng.Digest(i)
	}

	return true
}

func refresh(i *IndexController, h tango.Hook) {
	if resp, err := http.Get("/api/todo"); err == nil {
		defer resp.Body.Close()
		// load all todos from api-backend
		json.NewDecoder(resp.Body).Decode(&i.Todos)
		//time.Sleep(time.Second * 1)
		i.Busy = false
		h.Digest(i)
	}
}

func (i IndexController) Render() string {
	return tmplIndex
}

func (i *IndexController) AfterRender(tng tango.Hook) bool {
	i.Busy = true
	go refresh(i, tng)
	return true
}
