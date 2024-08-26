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

func (s *TodoStore) Insert(description string) (Todo, error) {
	stmt := `INSERT into todos (description) VALUES (?) RETURNING id, created_at, description, is_done`

    var t Todo

	err := s.DB.QueryRow(stmt, description).Scan(&t.ID, &t.CreatedAt, &t.Description, &t.IsDone)

	if err != nil {
		return Todo{}, err
	}

	return t, nil
}
