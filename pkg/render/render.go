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
func AddData(w http.ResponseWriter, tmpl string, td *models.TemplateData) *models.TemplateData {
	return td
}

// PageRender renders a template
func PageRender(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// get the template cache from the app config or from disk based on UseCache bool
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
		fmt.Println("Template was pulled from the cache")
	} else {
		tc, _ = CreateCache()
		fmt.Println("Template was pulled from the disk")
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get the template cache!")
	}
	// this causes the page to be created

	buf := new(bytes.Buffer)
	td = AddData(w, tmpl, td)
	_ = t.Execute(buf, td) // this is critical as this sends everything to the template, td contains the data
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

// CreateCache creates a template cache as a map
func CreateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
