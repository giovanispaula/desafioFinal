package dentista

import (
	"desafioII/internal/domain"
	"desafioII/pkg/store"
)

type Repository interface {
	Post(d domain.Dentista) (domain.Dentista, error)
	GetById(id int) (domain.Dentista, error)
	Update(id int, d domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
}

type repository struct {
	store store.DentistaStore
}

func NewRepository(store store.DentistaStore) Repository {
	return &repository{store}
}

func (r *repository) Post(d domain.Dentista) (domain.Dentista, error) {
	return r.store.Post(d)
}

func (r *repository) GetById(id int) (domain.Dentista, error) {
	return r.store.GetById(id)
}

func (r *repository) Update(id int, d domain.Dentista) (domain.Dentista, error) {
	return r.store.Update(id, d)
}

func (r *repository) Delete(id int) error {
	return r.store.Delete(id)
}
