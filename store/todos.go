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

func (s *TodoStore) Switch(id int) error {
	stmt := `SELECT is_done FROM todos WHERE id = ?`

	var is_Done bool
	err := s.DB.QueryRow(stmt, id).Scan(&is_Done)
	if err != nil {
		return err
	}
	newValue := !is_Done

	stmt2 := `UPDATE todos SET is_done = ? WHERE id = ?`
	_, err = s.DB.Exec(stmt2, newValue, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *TodoStore) Delete(id int) error {
	stmt := `DELETE FROM todos WHERE id = ?`

	_, err := s.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}
