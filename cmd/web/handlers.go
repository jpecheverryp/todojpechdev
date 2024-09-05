package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {

	todos, err := app.store.Todo.GetAll()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	files := []string{
		"./views/html/layout.html",
		"./views/html/index.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := TemplateData{
		Todos: todos,
	}
	err = ts.ExecuteTemplate(w, "layout", data)
	if err != nil {
		app.serverError(w, r, err)
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
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	files := []string{
		"./views/html/index.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := TemplateData{
		Todo: todo,
	}
	err = ts.ExecuteTemplate(w, "todo-component", data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) switchTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = app.store.Todo.Switch(id)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (app *application) deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = app.store.Todo.Delete(id)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (app *application) getAbout(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./views/html/layout.html",
		"./views/html/about.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}
