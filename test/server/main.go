package main

import (
	"net/http"

	"github.com/reinbach/golang-webapp-guide/test/template"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

var (
	templates = []string{"base.html"}
)

func Home(c web.C, w http.ResponseWriter, r *http.Request) {
	ctx := template.NewContext()
	ctx.Add("HomePage", true)
	template.Render(c, w, r, append(templates, "home.html"), ctx)
}

func About(c web.C, w http.ResponseWriter, r *http.Request) {
	ctx := template.NewContext()
	ctx.Add("AboutPage", true)
	template.Render(c, w, r, append(templates, "about.html"), ctx)
}

func NotFound(c web.C, w http.ResponseWriter, r *http.Request) {
	template.Render(c, w, r, append(templates, "404.html"),
		template.NewContext())
}

func main() {
	http.HandleFunc(template.STATIC_URL, template.StaticHandler)
	goji.Get("/", Home)
	goji.Get("/about", About)
	goji.NotFound(NotFound)

	goji.Serve()
}
