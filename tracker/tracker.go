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

func (t *Tracker) Add(name string, dueDate time.Time, priority Priority) (int, error) {
	r := Record{
		Name:          name,
		CreatedDate:   time.Now(),
		CompletedDate: time.Time{},
		DueDate:       dueDate,
		Priority:      priority,
	}

	if err := t.repo.CreateRecord(r); err != nil {
		return 0, err
	}

	t.records = append(t.records, r)
	t.sortRecords()

	index := -1
	for i, record := range t.records {
		if record.CreatedDate == r.CreatedDate {
			index = i
			break
		}
	}

	return index, nil
}

func (t *Tracker) sortRecords() {
	sort.Slice(t.records, func(i, j int) bool {
		return t.records[i].Urgency() > t.records[j].Urgency()
	})
}
