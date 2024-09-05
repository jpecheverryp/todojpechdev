package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"

	"todo.jpech.dev/store"
	"todo.jpech.dev/views"
)

type TemplateData struct {
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

	return cache, nil
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data any) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)

	err := ts.ExecuteTemplate(w, "layout", data)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

}
