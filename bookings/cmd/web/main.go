package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/YanAmorelli/bookings/internal/config"
	"github.com/YanAmorelli/bookings/internal/handlers"
	"github.com/YanAmorelli/bookings/internal/helpers"
	"github.com/YanAmorelli/bookings/internal/models"
	"github.com/YanAmorelli/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager
var errorLog *log.Logger
var infoLog *log.Logger

func main() {
	gob.Register(models.Reservation{})

	// change this to true when is in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "Info: \t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "Error: \t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog



	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)
	helpers.NewHelpers(&app)
	
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