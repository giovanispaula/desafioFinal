package paciente

import (
	"desafioII/internal/domain"
	"desafioII/pkg/store"
)

type Repository interface {
	Post(p domain.Paciente) (domain.Paciente, error)
	GetById(id int) (domain.Paciente, error)
	Update(id int, p domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
}

type repository struct {
	store store.PacienteStore
}

func NewRepository(store store.PacienteStore) Repository {
	return &repository{store}
}

func (r *repository) Post(p domain.Paciente) (domain.Paciente, error) {
	return r.store.Post(p)
}

func (r *repository) GetById(id int) (domain.Paciente, error) {
	return r.store.GetById(id)
}

func (r *repository) Update(id int, p domain.Paciente) (domain.Paciente, error) {
	return r.store.Update(id, p)
}

func (r *repository) Delete(id int) error {
	return r.store.Delete(id)
}
