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

func (s *TodoStore) GetAll() ([]Todo, error) {
	stmt := `SELECT id, created_at, description, is_done FROM todos ORDER BY created_at DESC`

	var todos = []Todo{}

	rows, err := s.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var t Todo

		err := rows.Scan(&t.ID, &t.CreatedAt, &t.Description, &t.IsDone)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}
