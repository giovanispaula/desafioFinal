package store

import (
	"database/sql"
	"desafioII/internal/domain"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type ConsultaStore interface {
	Post(c domain.Consulta) (domain.ConsultaDTO, error)
	GetById(id int) (domain.ConsultaDTO, error)
	Update(id int, c domain.Consulta) (domain.ConsultaDTO, error)
	Delete(id int) error
}

func NewConsultaStore() *consultaStore {
	dbConsulta, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	return &consultaStore{
		db: dbConsulta,
	}
}

type consultaStore struct {
	db *sql.DB
}

func (cs *consultaStore) Post(c domain.Consulta) (domain.ConsultaDTO, error) {
	var consulta domain.ConsultaDTO

	result, err := cs.db.Exec("INSERT INTO consultas (descricao, dataConsulta, dentistaId, pacienteId) VALUES (?,?,?,?)", c.Descricao, c.DataConsulta, c.DentistaId, c.PacienteId)

	if err != nil {
		return domain.ConsultaDTO{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return domain.ConsultaDTO{}, err
	}

	consulta.Id = int(id)

	return cs.GetById(consulta.Id)
}

func (cs *consultaStore) GetById(id int) (domain.ConsultaDTO, error) {
	var c domain.ConsultaDTO

	result, err := cs.db.Query("SELECT c.id, c.descricao, c.dataConsulta, c.dentistaId, c.pacienteId, d.id, d.nome, d.sobrenome, d.matricula, p.id, p.nome, p.sobrenome, p.rg, p.dataCadastro FROM consultas c inner JOIN dentistas d on c.dentistaId = d.id inner JOIN pacientes p on c.pacienteId = p.id WHERE c.id=?", id)

	defer result.Close()

	if err != nil {
		return c, err
	}

	for result.Next() {
		if err := result.Scan(
			&c.Id,
			&c.Descricao,
			&c.DataConsulta,
			&c.DentistaId,
			&c.PacienteId,
			&c.Dentista.Id,
			&c.Dentista.Nome,
			&c.Dentista.Sobrenome,
			&c.Dentista.Matricula,
			&c.Paciente.Id,
			&c.Paciente.Nome,
			&c.Paciente.Sobrenome,
			&c.Paciente.Rg,
			&c.Paciente.DataCadastro); err != nil {
			return c, err
		}

		return c, nil
	}

	if result.Next() {
		return c, nil
	}

	return domain.ConsultaDTO{}, errors.New("Consulta não registrada")
}

func (cs *consultaStore) Update(id int, c domain.Consulta) (domain.ConsultaDTO, error) {

	_, err := cs.db.Exec("UPDATE consultas SET descricao=?, dataConsulta=?, dentistaId=?, pacienteId=? WHERE id=?", c.Descricao, c.DataConsulta, c.DentistaId, c.PacienteId, id)

	if err != nil {
		return domain.ConsultaDTO{}, err
	}

	return cs.GetById(id)
}

func (cs *consultaStore) Delete(id int) error {

	result, err := cs.db.Exec("DELETE FROM consultas WHERE id = ?", id)

	if err != nil {
		return err
	}

	count, err := result.RowsAffected()

	if err != nil {
		return errors.New("Consulta não registrada")
	}

	if count != 0 {
		return nil
	}

	return err
}
