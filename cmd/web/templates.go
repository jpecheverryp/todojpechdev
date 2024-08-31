package main

import "todo.jpech.dev/store"

type TemplateData struct {
	Todos []store.Todo
	Todo  store.Todo
}
