package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/YanAmorelli/bookings/pkg/config"
	"github.com/YanAmorelli/bookings/pkg/handlers"
	"github.com/YanAmorelli/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	// change this to true when is in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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