package main

import (
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

	db, err := openDB("./data/sqlite.db")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger: logger,
		store:  store.NewStore(db),
	}

	logger.Info("starting server", "port", port)
	err = http.ListenAndServe(port, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
