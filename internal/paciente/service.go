package paciente

import "desafioII/internal/domain"

type Service interface {
	Post(p domain.Paciente) (domain.Paciente, error)
	GetById(id int) (domain.Paciente, error)
	Update(id int, p domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Post(p domain.Paciente) (domain.Paciente, error) {
	return s.r.Post(p)
}

func (s *service) GetById(id int) (domain.Paciente, error) {
	return s.r.GetById(id)
}

func (s *service) Update(id int, p domain.Paciente) (domain.Paciente, error) {
	paciente, err := s.r.GetById(id)

	if err != nil {
		return domain.Paciente{}, err
	}

	if p.Nome == "" {
		p.Nome = paciente.Nome
	}

	if p.Sobrenome == "" {
		p.Sobrenome = paciente.Sobrenome
	}

	if p.Rg == "" {
		p.Rg = paciente.Rg
	}

	if p.DataCadastro == "" {
		p.DataCadastro = paciente.DataCadastro
	}

	return s.r.Update(id, p)
}

func (s *service) Delete(id int) error {
	return s.r.Delete(id)
}
