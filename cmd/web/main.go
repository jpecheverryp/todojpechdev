package main

import (
	"log"
	"net/http"
)

type application struct {
    
}

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello Todo"))
}

func main() {
    port := ":5174"
    mux := http.NewServeMux()

    app := &application{}

    mux.HandleFunc("/", app.getIndex)

    log.Printf("starting server in port: %s", port)
    log.Fatal(http.ListenAndServe(port, mux))
}
