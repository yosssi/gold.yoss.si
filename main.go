// Package main provides main processes.
package main

import (
	"fmt"
	"net/http"

	"github.com/martini-contrib/render"
	"github.com/yosssi/rendergold"
	"github.com/yosssi/staticbin"
)

var (
	stringTemplate = map[string]string{}
	templateNames  = []string{
		"layout",
		"top",
	}
)

// main executes main processes.
func main() {
	setStringTemplate()
	m := staticbin.Classic(Asset)
	m.Use(rendergold.Renderer())
	m.Get("/", func(r render.Render) {
		r.HTML(
			http.StatusOK,
			"top",
			map[string]interface{}{
				"Title": "Gold - Template engine for Go",
			},
			render.HTMLOptions{
				Layout: "layout" + rendergold.NameContentDelim + stringTemplate["layout"],
			},
			render.HTMLOptions{
				Layout: "top" + rendergold.NameContentDelim + stringTemplate["top"],
			},
		)
	})
	m.Run()
}

// setStringTemplate sets string template.
func setStringTemplate() {
	for _, name := range templateNames {
		b, err := Asset(fmt.Sprintf("templates/%s.gold", name))
		if err != nil {
			panic(err)
		}
		stringTemplate[name] = string(b)
	}
}
