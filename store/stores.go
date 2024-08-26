package store

import "database/sql"

type Store struct {
	Todo TodoStore
}

func NewStore(db *sql.DB) Store {
	return Store{
		Todo: TodoStore{DB: db},
	}
}
