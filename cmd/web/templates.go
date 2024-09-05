package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"

	"todo.jpech.dev/store"
	"todo.jpech.dev/views"
)

type templateData struct {
	Todos []store.Todo
	Todo  store.Todo
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(views.Files, "html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/layout.html",
			page,
		}

		ts, err := template.New(name).ParseFS(views.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	components, err := fs.Glob(views.Files, "html/components/*.html")
	if err != nil {
		return nil, err
	}

	for _, component := range components {
		name := filepath.Base(component)

		ts, err := template.New(name).ParseFS(views.Files, component)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

func (app *application) newTemplateData() templateData {
	return templateData{}
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, r, err)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "layout", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}


func (app *application) renderComponent(w http.ResponseWriter, r *http.Request, status int, component string, data templateData) {
	ts, ok := app.templateCache[component]
	if !ok {
		err := fmt.Errorf("the component %s does not exist", component)
		app.serverError(w, r, err)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, component, data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}
