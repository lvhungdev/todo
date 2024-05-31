package storage

import (
	"database/sql"

	"github.com/lvhungdev/todo/tracker"
	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db *sql.DB
}

func NewStore(path string) (*Store, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	store := &Store{db: db}
	if err = store.migrate(); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *Store) Close() {
	if err := s.db.Close(); err != nil {
		// TODO handle error
	}
}

func (s *Store) GetAllActiveRecords() ([]tracker.Record, error) {
	rows, err := s.db.Query("SELECT id, name, created_date, completed_date, due_date, priority FROM records WHERE completed_date = '0001-01-01 00:00:00+00:00'")
	if err != nil {
		return nil, err
	}

	records := []tracker.Record{}
	for rows.Next() {
		var r tracker.Record
		if err := rows.Scan(&r.Id, &r.Name, &r.CreatedDate, &r.CompletedDate, &r.DueDate, &r.Priority); err != nil {
			return nil, err
		}
		records = append(records, r)
	}

	return records, nil
}

func (s *Store) CreateRecord(r tracker.Record) error {
	_, err := s.db.Exec(
		"INSERT INTO records(id, name, created_date, completed_date, due_date, priority) VALUES (?, ?, ?, ?, ?, ?)",
		r.Id,
		r.Name,
		r.CreatedDate,
		r.CompletedDate,
		r.DueDate,
		r.Priority,
	)
	return err
}

func (s *Store) UpdateRecord(r tracker.Record) error {
	_, err := s.db.Exec(
		"UPDATE records SET name = ?, created_date = ?, completed_date = ?, due_date = ?, priority = ? WHERE id = ?",
		r.Name,
		r.CreatedDate,
		r.CompletedDate,
		r.DueDate,
		r.Priority,
		r.Id,
	)
	return err
}

func (s *Store) migrate() error {
	_, err := s.db.Exec(
		`CREATE TABLE IF NOT EXISTS records (
    		id TEXT PRIMARY KEY NOT NULL,
    		'name' TEXT NOT NULL,
    		created_date TIMESTAMP NOT NULL,
    		completed_date TIMESTAMP NOT NULL,
    		due_date TIMESTAMP NOT NULL,
    		priority INTEGER NOT NULL
		)`)
	return err
}
