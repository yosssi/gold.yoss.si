// Package main provides main processes.
package main

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/yosssi/rendergold"
)

// main executes main processes.
func main() {
	m := martini.Classic()
	m.Use(rendergold.Renderer())
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
