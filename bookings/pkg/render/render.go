package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	
	"github.com/YanAmorelli/bookings/pkg/config"
	"github.com/YanAmorelli/bookings/pkg/models"
)

var functions = template.FuncMap{
}

var app *config.AppConfig
func NewTemplates(a *config.AppConfig){
	app = a
}

func AddDefaultData(data *models.TemplateData) *models.TemplateData {
	return data
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData) {
	var templateCache map[string]*template.Template
	
	if app.UseCache {
		templateCache = app.TemplateCache
		} else {
			templateCache, _ = CreateTemplateCache()
		}
		
		parsedTemplate, ok := templateCache[tmpl]
		if !ok {
			log.Fatal("Couldn't get template from template cache")
		}
		
		buffer := new(bytes.Buffer)

		data = AddDefaultData(data)

		_ = parsedTemplate.Execute(buffer, data)
		
		_, err := buffer.WriteTo(w)
		if err != nil {
			fmt.Println("Error writing template to browser", err)
		}
		
	}
	
	func CreateTemplateCache() (map[string]*template.Template, error) {
		
		myCache := map[string]*template.Template{}
		
		pages, err := filepath.Glob("./templates/*.page.go.tmpl")
		if err != nil {
			return myCache, err
		}
		
		for _, page := range pages {
			name := filepath.Base(page)
			
			templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
			if err != nil {
				return myCache, err
			}
			
			matchLayout, err := filepath.Glob("./templates/*.layout.go.tmpl")
			if err != nil {
				return myCache, err
			}
			
			if len(matchLayout) > 0 {
				templateSet, err = templateSet.ParseGlob("./templates/*.layout.go.tmpl")
				if err != nil {
					return myCache, err
				}
			}
			
			myCache[name] = templateSet
			
		}
		return myCache, err
	}