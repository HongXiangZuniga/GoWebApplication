package handlers

import (
	"net/http"

	"github.com/HongXiangZuniga/GoWebApplication/pkg/config"
	"github.com/HongXiangZuniga/GoWebApplication/pkg/render"
	"github.com/HongXiangZuniga/GoWebApplication/templates"
)

//TemplateData holds data sents from handlers to template

// Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

func NewRepo(appConfig *config.AppConfig) *Repository {
	return &Repository{
		App: appConfig}
}

//New Handler sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &templates.TemplateData{})
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	render.RenderTemplate(w, "about.page.tmpl", &templates.TemplateData{
		StringMap: stringMap,
	})
}
