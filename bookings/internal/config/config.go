package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	// UseCache is a configuration with the idea of development and production environments.
	// If it's on dev then don't use cache because it's faster to see changes. 
	UseCache		bool
	TemplateCache 	map[string]*template.Template
	InfoLog 		*log.Logger
	ErrorLog 		*log.Logger
	InProduction 	bool
	Session 			*scs.SessionManager
}