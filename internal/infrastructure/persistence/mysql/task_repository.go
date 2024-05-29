package mysql

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"

	"task-api/internal/domain/task"
)

type TaskRepository struct {
	DB *sql.DB
}

func NewMySQLTaskRepository(dsn string) (task.Repository, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &TaskRepository{DB: db}, nil
}

func (r *TaskRepository) GetAll() ([]task.Info, error) {
	rows, err := r.DB.Query("SELECT id, name, status FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []task.Info
	for rows.Next() {
		var t task.Info
		if err := rows.Scan(&t.ID, &t.Name, &t.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *TaskRepository) GetByID(id int) (task.Info, error) {
	var t task.Info
	err := r.DB.QueryRow("SELECT id, name, status FROM tasks WHERE id = ?", id).Scan(&t.ID, &t.Name, &t.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return task.Info{}, errors.New("task not found")
		}
		return task.Info{}, err
	}
	return t, nil
}

func (r *TaskRepository) Create(t task.Info) (task.Info, error) {
	result, err := r.DB.Exec("INSERT INTO tasks (name, status) VALUES (?, ?)", t.Name, t.Status)
	if err != nil {
		return task.Info{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return task.Info{}, err
	}
	t.ID = int(id)
	return t, nil
}

func (r *TaskRepository) Update(t task.Info) (task.Info, error) {
	_, err := r.DB.Exec("UPDATE tasks SET name = ?, status = ? WHERE id = ?", t.Name, t.Status, t.ID)
	if err != nil {
		return task.Info{}, err
	}
	return t, nil
}

func (r *TaskRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
