package tracker

import (
	"sort"
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

	t := &Tracker{
		repo:    repo,
		records: records,
	}
	t.sortRecords()

	return t, nil
}

func (t *Tracker) ListActive() []Record {
	return t.records
}

func (t *Tracker) Add(name string, dueDate time.Time, priority Priority) error {
	r := Record{
		Name:          name,
		CreatedDate:   time.Now(),
		CompletedDate: time.Time{},
		DueDate:       dueDate,
		Priority:      priority,
	}

	if err := t.repo.CreateRecord(r); err != nil {
		return err
	}

	t.records = append(t.records, r)
	t.sortRecords()

	return nil
}

func (t *Tracker) sortRecords() {
	sort.Slice(t.records, func(i, j int) bool {
		return t.records[i].Urgency() > t.records[j].Urgency()
	})
}
