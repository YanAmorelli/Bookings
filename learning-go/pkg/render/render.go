package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	templateCache, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	parsedTemplate, ok := templateCache[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buffer := new(bytes.Buffer)

	_ = parsedTemplate.Execute(buffer, nil)

	_, err = buffer.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

func createTemplateCache() (map[string]*template.Template, error) {
	
	myCache := map[string]*template.Template{}
	
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	
	for _, page := range pages {
		name := filepath.Base(page)
		
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		
		matchLayout, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matchLayout) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet

	}
	return myCache, err
}