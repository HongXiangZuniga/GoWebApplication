package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HongXiangZuniga/GoWebApplication/pkg/config"
	"github.com/HongXiangZuniga/GoWebApplication/pkg/handlers"
	"github.com/HongXiangZuniga/GoWebApplication/pkg/render"
)

const portNumber = ":8081"

func main() {
	var app config.AppConfig

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
