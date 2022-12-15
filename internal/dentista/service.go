package dentista

import "desafioII/internal/domain"

type Service interface {
	Post(d domain.Dentista) (domain.Dentista, error)
	GetById(id int) (domain.Dentista, error)
	Update(id int, d domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Post(d domain.Dentista) (domain.Dentista, error) {
	return s.r.Post(d)
}

func (s *service) GetById(id int) (domain.Dentista, error) {
	return s.r.GetById(id)
}

func (s *service) Update(id int, d domain.Dentista) (domain.Dentista, error) {
	dentista, err := s.r.GetById(id)

	if err != nil {
		return domain.Dentista{}, err
	}

	if d.Nome == "" {
		d.Nome = dentista.Nome
	}

	if d.Sobrenome == "" {
		d.Sobrenome = dentista.Sobrenome
	}

	if d.Matricula == "" {
		d.Matricula = dentista.Matricula
	}

	return s.r.Update(id, d)
}

func (s *service) Delete(id int) error {
	return s.r.Delete(id)
}
