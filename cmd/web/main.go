package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"todo.jpech.dev/store"
)

type application struct {
	logger *slog.Logger
	store  store.Store
}

func main() {
	port := ":5174"

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mux := http.NewServeMux()

	db, err := openDB("./data/sqlite.db")
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		logger: logger,
		store:  store.NewStore(db),
	}

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	//mux.Handle("GET /static/", http.FileServerFS(static.Files))

	mux.HandleFunc("GET /", app.getIndex)
	mux.HandleFunc("POST /todo", app.createTodo)
	mux.HandleFunc("PUT /switch-todo/{id}", app.switchTodo)
	mux.HandleFunc("DELETE /todo/{id}", app.deleteTodo)

	logger.Info("starting server", "port", port)
	err = http.ListenAndServe(port, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
