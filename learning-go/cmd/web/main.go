package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/YanAmorelli/learning-go/pkg/config"
	"github.com/YanAmorelli/learning-go/pkg/handlers"
	"github.com/YanAmorelli/learning-go/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)
	
	// Setting environment to dev 
	app.UseCache = false
	
	// Setting environment to production
	// app.UseCache = true
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = templateCache
	render.NewTemplates(&app)
	
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	
	serve := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}