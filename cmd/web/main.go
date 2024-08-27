package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
	"todo.jpech.dev/store"
)

type application struct {
	Store store.Store
}

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {

	todos, err := app.Store.Todo.GetAll()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	files := []string{
		"./views/index.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "index", todos)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (app *application) createTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	newTodoDescription := r.PostForm.Get("new-todo")
	todo, err := app.Store.Todo.Insert(newTodoDescription)
	log.Print(todo)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
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
	mux.HandleFunc("POST /todo", app.createTodo)

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
