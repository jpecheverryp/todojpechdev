package main

import (
	"net/http"

	"todo.jpech.dev/views"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(views.Files))

	mux.HandleFunc("GET /", app.getIndex)
	mux.HandleFunc("POST /todo", app.createTodo)
	mux.HandleFunc("PUT /switch-todo/{id}", app.switchTodo)
	mux.HandleFunc("DELETE /todo/{id}", app.deleteTodo)

	mux.HandleFunc("GET /about", app.getAbout)

	return mux
}
