package main

import (
	"database/sql"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	_ "modernc.org/sqlite"
	"todo.jpech.dev/store"
)

type application struct {
    logger *slog.Logger
	store store.Store
}

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {

	todos, err := app.store.Todo.GetAll()
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
	todo, err := app.store.Todo.Insert(newTodoDescription)
	log.Print(todo)
	if err != nil {
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

	err = ts.ExecuteTemplate(w, "todo", todo)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (app *application) switchTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = app.store.Todo.Switch(id)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (app *application) deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = app.store.Todo.Delete(id)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
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
		store: store.NewStore(db),
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

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
