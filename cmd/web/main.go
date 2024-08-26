package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
	"todo.jpech.dev/store"
)

type application struct {
	Store store.Store
}

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Todo"))
}

func main() {
	port := ":5174"
	mux := http.NewServeMux()

	db, err := openDB("./data/sqlite.db")
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		Store: store.NewStore(db),
	}

	mux.HandleFunc("/", app.getIndex)

	log.Printf("starting server in port: %s", port)
	log.Fatal(http.ListenAndServe(port, mux))
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
