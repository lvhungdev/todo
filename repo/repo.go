package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/lvhungdev/todo/core"
	_ "github.com/mattn/go-sqlite3"
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo() *Repo {
	return &Repo{
		db: nil,
	}
}

func (r *Repo) Init(driverName string, dataSourceName string) error {
	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return err
	}

	r.db = db

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS todos
    (
        id             INTEGER PRIMARY KEY,
        name           TEXT NOT NULL,
        created_date   DATETIME NOT NULL,
        completed_date DATETIME,
        is_completed   INTEGER NOT NULL,
        due_date       DATETIME,
        priority       INTEGER
    )
    `)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) Get() ([]core.Todo, error) {
	todos := []core.Todo{}
	err := r.db.Select(&todos, "SELECT * FROM todos")

	return todos, err
}

func (r *Repo) GetById(id int) (core.Todo, error) {
	todo := core.Todo{}
	err := r.db.Get(&todo, "SELECT * FROM todos WHERE id=$1", id)

	return todo, err
}

func (r *Repo) Create(todo *core.Todo) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	result, err := tx.Exec(`
    INSERT INTO todos
    (name, created_date, completed_date, is_completed, due_date, priority)
    VALUES
    ($1, $2, $3, $4, $5, $6)
    `, todo.Name, todo.CreatedDate, todo.CompletedDate, todo.IsCompleted, todo.DueDate, todo.Priority)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	todo.Id = int(id)

	err = tx.Commit()
	if err != nil {
		return err
	}

	return err
}
