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

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = templateCache
	render.NewTemplates(&app)

	app.UseCache = false

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
