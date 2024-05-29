package tracker

import (
	"time"
)

type Repo interface {
	GetAllActiveRecords() ([]Record, error)
	CreateRecord(Record) error
}

type Tracker struct {
	repo    Repo
	records []Record
}

func NewTracker(repo Repo) (*Tracker, error) {
	records, err := repo.GetAllActiveRecords()
	if err != nil {
		return nil, err
	}

	return &Tracker{
		repo:    repo,
		records: records,
	}, nil
}

func (t *Tracker) ListActive() []Record {
	return t.records
}

func (t *Tracker) Add(name string, dueDate time.Time) error {
	r := Record{
		Name:          name,
		CreatedDate:   time.Now(),
		CompletedDate: time.Time{},
		DueDate:       dueDate,
		Priority:      PriNone,
	}

	if err := t.repo.CreateRecord(r); err != nil {
		return err
	}

	t.records = append(t.records, r)
	return nil
}
