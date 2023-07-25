package render

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	// "github.com/supWRLD/bookings/pckg/config"
	"github.com/supWRLD/bookings/pckg/models"
)

// var app *config.AppConfig
// app = CreateTemplateCache()

// func NewTemplates(a *config.AppConfig) {
// 	app = a
// }

func AddDefaultData(tmplData *models.TemplateData) *models.TemplateData {
	return tmplData
}

func RenderTemplate(w http.ResponseWriter, tmpl string, tmplData *models.TemplateData) {
	var tmplCache map[string]*template.Template

	tmplCache, _ = CreateTemplateCache()

	// get requested template from cache
	t, ok := tmplCache[tmpl]
	if !ok {
		log.Fatal("couldn't get template from template cache")
	}

	tmplData = AddDefaultData(tmplData)
	// render the template
	err := t.Execute(w, tmplData)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// get the template cache from main
	myCache := map[string]*template.Template{}

	// get all the files named *.html from templates folder
	pages, err := filepath.Glob("../../templates/*.html")
	if err != nil {
		return myCache, err
	}

	// go through the pages slice
	for _, page := range pages {
		// get the .html name only
		name := filepath.Base(page)

		// creating template
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		myCache[name] = templateSet
	}
	return myCache, nil
}
