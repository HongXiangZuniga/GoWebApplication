package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/HongXiangZuniga/GoWebApplication/pkg/config"
	"github.com/HongXiangZuniga/GoWebApplication/templates"
)

var functions = template.FuncMap{}
var app *config.AppConfig

//New Templates sets the config for the template package
func NewTemplates(appConfig *config.AppConfig) {
	app = appConfig
}

func AddDefaultData(templateData *templates.TemplateData) *templates.TemplateData {

	return templateData
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data *templates.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	//get the template cache from the app conff
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(("Could not get template from template cache"))
	}
	buf := new(bytes.Buffer)
	templateData := AddDefaultData(data)
	_ = t.Execute(buf, templateData)
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to:" + err.Error())
	}
}

//Create template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myChache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myChache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myChache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myChache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myChache, err
			}
		}
		myChache[name] = ts
	}
	return myChache, nil
}
