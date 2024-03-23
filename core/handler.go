package core

import "errors"

type Repo interface {
	Get() ([]Todo, error)
	GetById(id int) (Todo, error)
	Create(todo *Todo) error
	// Update(todo *Todo) error
	// Delete(id int) error
}

type Handler struct {
	repo  Repo
	todos []Todo
}

func NewHandler(repo Repo) *Handler {
	return &Handler{
		repo:  repo,
		todos: []Todo{},
	}
}

func (h *Handler) List() ([]Todo, error) {
	return h.repo.Get()
}

func (h *Handler) Add(name string) (Todo, error) {
	if len(name) == 0 {
		return Todo{}, errors.New("[ERR] name is required")
	}

	todo := &Todo{
		name: name,
	}

	err := h.repo.Create(todo)
	if err != nil {
		return Todo{}, err
	}

	return *todo, nil
}
