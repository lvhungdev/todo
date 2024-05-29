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

	return &Store{db: db}, nil
}

func (s *Store) Close() {
	if err := s.db.Close(); err != nil {
		// TODO handle error
	}
}

func (s *Store) GetAllActiveRecords() ([]tracker.Record, error) {
	rows, err := s.db.Query("SELECT name, created_date, completed_date, due_date, priority FROM records WHERE completed_date = '0001-01-01 00:00:00+00:00'")
	if err != nil {
		return nil, err
	}

	records := []tracker.Record{}
	for rows.Next() {
		var r tracker.Record
		if err := rows.Scan(&r.Name, &r.CreatedDate, &r.CompletedDate, &r.DueDate, &r.Priority); err != nil {
			return nil, err
		}
		records = append(records, r)
	}

	return records, nil
}

func (s *Store) CreateRecord(r tracker.Record) error {
	stmt, err := s.db.Prepare("INSERT INTO records(name, created_date, completed_date, due_date, priority) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.Name, r.CreatedDate, r.CompletedDate, r.DueDate, r.Priority)
	if err != nil {
		return err
	}

	return nil
}
