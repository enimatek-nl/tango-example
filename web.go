package main

import (
	"github.com/enimatek-nl/tango"
	"github.com/enimatek-nl/tango-example/web"
	"github.com/enimatek-nl/tango/std"
)

func main() {
	tg := tango.New()

	tg.AddComponents(
		std.Router{},
		std.Repeat{},
		std.Click{},
		std.Bind{},
		std.Change{},
		std.Model{},
		std.Attr{},
		web.Busy{})

	tg.AddRoutes(
		tango.NewRoute("/", &web.IndexController{}),
		tango.NewRoute("/edit/:id", &web.EditController{}),
	)

	tg.Bootstrap()

	<-make(chan bool)
}
