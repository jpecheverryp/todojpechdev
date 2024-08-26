package store

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID          int64
	CreatedAt   time.Time
	Description string
	IsDone      bool
}

type TodoStore struct {
	DB *sql.DB
}
