package store

import (
	"database/sql"
	"desafioII/internal/domain"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type PacienteStore interface {
	Post(p domain.Paciente) (domain.Paciente, error)
	GetById(id int) (domain.Paciente, error)
	Update(id int, p domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
}

func NewPacienteStore() PacienteStore {
	dbPaciente, err := ConnectDb()

	if err != nil {
		panic(err)
	}

	return &pacienteStore{
		db: dbPaciente,
	}
}

type pacienteStore struct {
	db *sql.DB
}

func (ps *pacienteStore) Post(p domain.Paciente) (domain.Paciente, error) {
	var paciente domain.Paciente

	result, err := ps.db.Exec("INSERT INTO pacientes (nome, sobrenome, rg, dataCadastro) VALUES (?,?,?,?)", p.Nome, p.Sobrenome, p.Rg, p.DataCadastro)

	if err != nil {
		return paciente, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return paciente, err
	}

	paciente.Id = int(id)

	return ps.GetById(paciente.Id)
}

func (ps *pacienteStore) GetById(id int) (domain.Paciente, error) {
	var p domain.Paciente

	result, err := ps.db.Query("SELECT * FROM pacientes WHERE id = ?", id)

	defer result.Close()

	if err != nil {
		return domain.Paciente{}, err
	}

	defer result.Close()

	for result.Next() {
		if err = result.Scan(
			&p.Id,
			&p.Nome,
			&p.Sobrenome,
			&p.Rg,
			&p.DataCadastro); err != nil {
			return domain.Paciente{}, err
		}
		return p, nil
	}
	if result.Next() {
		return p, nil
	}
	return domain.Paciente{}, errors.New("Paciente não cadastrado")
}

func (ps *pacienteStore) Update(id int, p domain.Paciente) (domain.Paciente, error) {

	_, err := ps.db.Exec("UPDATE pacientes SET nome=?, sobrenome=?, rg=?, data_cadastro=? WHERE id=?", p.Nome, p.Sobrenome, p.Rg, p.DataCadastro, id)

	if err != nil {
		return p, err
	}

	return ps.GetById(id)
}

func (ps *pacienteStore) Delete(id int) error {

	result, err := ps.db.Exec("DELETE FROM pacientes WHERE id = ?", id)

	if err != nil {
		return err
	}

	count, err := result.RowsAffected()

	if err != nil {
		return errors.New("Paciente não cadastrado")
	}

	if count != 0 {
		return nil
	}

	return err
}
