package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	//mux.Handle("GET /static/", http.FileServerFS(static.Files))

	mux.HandleFunc("GET /", app.getIndex)
	mux.HandleFunc("POST /todo", app.createTodo)
	mux.HandleFunc("PUT /switch-todo/{id}", app.switchTodo)
	mux.HandleFunc("DELETE /todo/{id}", app.deleteTodo)

	return mux
}
