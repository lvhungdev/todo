package tracker

import (
	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"
)

type Repo interface {
	GetAllActiveRecords() ([]Record, error)
	CreateRecord(Record) error
	UpdateRecord(Record) error
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

func (t *Tracker) GetActive() []Record {
	result := []Record{}
	for _, r := range t.records {
		if !r.Completed() {
			result = append(result, r)
		}
	}

	return result
}

func (t *Tracker) Add(name string, dueDate time.Time, priority Priority) (int, error) {
	r := Record{
		Id:            uuid.NewString(),
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
		if record.Id == r.Id {
			index = i
			break
		}
	}

	return index, nil
}

func (t *Tracker) Complete(index int) error {
	if index < 0 || index >= len(t.records) {
		return fmt.Errorf("[ERR] invalid index: %d", index)
	}

	t.records[index].CompletedDate = time.Now()
	return t.repo.UpdateRecord(t.records[index])
}

func (t *Tracker) Modify(index int, name *string, dueDate *time.Time, priority *Priority) error {
	if index < 0 || index >= len(t.records) {
		return fmt.Errorf("[ERR] invalid index: %d", index)
	}

	if name != nil && *name != "" {
		t.records[index].Name = *name
	}
	if dueDate != nil {
		t.records[index].DueDate = *dueDate
	}
	if priority != nil {
		t.records[index].Priority = *priority
	}

	return t.repo.UpdateRecord(t.records[index])
}

func (t *Tracker) sortRecords() {
	sort.Slice(t.records, func(i, j int) bool {
		return t.records[i].Urgency() > t.records[j].Urgency()
	})
}
