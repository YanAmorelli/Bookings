package config

import "text/template"

type AppConfig struct {
	// UseCache is a configuration with the idea of development and production environments.
	// If it's on dev then don't use cache because it's faster to see changes. 
	UseCache		bool
	TemplateCache 	map[string]*template.Template
}