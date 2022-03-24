package handlers

import (
	"net/http"

	"github.com/HongXiangZuniga/GoWebApplication/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.go.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.go.tmpl")
}
