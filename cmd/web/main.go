package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/HongXiangZuniga/GoWebApplication/pkg/config"
	"github.com/HongXiangZuniga/GoWebApplication/pkg/handlers"
	"github.com/HongXiangZuniga/GoWebApplication/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8081"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction //false in localhost
	app.Session = session
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Panic("Not loading template Cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting app in port %s", portNumber))
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
