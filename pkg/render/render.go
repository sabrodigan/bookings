package render

import (
	"bytes"
	"fmt"
	"github.com/sabrodigan/bookings/models"
	"github.com/sabrodigan/bookings/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// New Templates renders templates using html/template
var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}
func AddData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// get the template cache from the app config or from disk based on UseCache bool
	var tc map[string]*template.Template
	
	config.WorkingDir()
	
	if app.UseCache {
		tc = app.TemplateCache
		fmt.Println("Template was pulled from the cache")
	} else {
		tc, _ = CreateTemplateCache()
		fmt.Println("Template was pulled from the disk")
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get the template cache!")
	}
	// this causes the page to be created
	buf := new(bytes.Buffer)
	td = AddData(td)
	_ = t.Execute(buf, td) // this is critical as this sends everything to the template, td contains the data
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	fmt.Println("Creating template cache")
	
	config.WorkingDir()

	myCache := map[string]*template.Template{}
	

	fmt.Println("app.InProduction", app.TemplateCache)
	
	pages, err := filepath.Glob("/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
