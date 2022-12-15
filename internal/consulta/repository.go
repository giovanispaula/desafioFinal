package consulta

import (
	"desafioII/internal/domain"
	"desafioII/pkg/store"
)

type Repository interface {
	Post(c domain.Consulta) (domain.ConsultaDTO, error)
	GetById(id int) (domain.ConsultaDTO, error)
	Update(id int, c domain.Consulta) (domain.ConsultaDTO, error)
	Delete(id int) error
}

type repository struct {
	store store.ConsultaStore
}

func NewRepository(store store.ConsultaStore) Repository {
	return &repository{store}
}

func (r *repository) Post(c domain.Consulta) (domain.ConsultaDTO, error) {
	return r.store.Post(c)
}

func (r *repository) GetById(id int) (domain.ConsultaDTO, error) {
	return r.store.GetById(id)
}

func (r *repository) Update(id int, c domain.Consulta) (domain.ConsultaDTO, error) {
	return r.store.Update(id, c)
}

func (r *repository) Delete(id int) error {
	return r.store.Delete(id)
}
