// Package main provides main processes.
package main

import (
	"net/http"

	"github.com/martini-contrib/render"
	"github.com/yosssi/rendergold"
	"github.com/yosssi/staticbin"
)

// main executes main processes.
func main() {
	m := staticbin.Classic(Asset)
	m.Use(rendergold.Renderer(rendergold.Options{Asset: Asset}))
	m.Get("/", func(r render.Render) {
		r.HTML(
			http.StatusOK,
			"top",
			map[string]interface{}{
				"Title": "Gold - Template engine for Go",
			},
		)
	})
	m.Run()
}
