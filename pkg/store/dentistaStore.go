package store

import (
	"database/sql"
	"desafioII/internal/domain"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type DentistaStore interface {
	Post(p domain.Dentista) (domain.Dentista, error)
	GetById(id int) (domain.Dentista, error)
	Update(id int, d domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
}

func NewDentistaStore() DentistaStore {
	dbDentista, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	return &dentistaStore{
		db: dbDentista,
	}
}

type dentistaStore struct {
	db *sql.DB
}

func (ds *dentistaStore) Post(d domain.Dentista) (domain.Dentista, error) {
	var dentista domain.Dentista

	result, err := ds.db.Exec("INSERT INTO dentistas (nome, sobrenome, matricula) VALUES (?,?,?)", d.Nome, d.Sobrenome, d.Matricula)

	if err != nil {
		return dentista, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return dentista, err
	}

	dentista.Id = int(id)

	return ds.GetById(dentista.Id)
}

func (ds *dentistaStore) GetById(id int) (domain.Dentista, error) {
	var d domain.Dentista

	result, err := ds.db.Query("SELECT * FROM dentistas WHERE id = ?", id)

	defer result.Close()

	if err != nil {
		return d, err
	}

	for result.Next() {
		if err := result.Scan(
			&d.Id,
			&d.Nome,
			&d.Sobrenome,
			&d.Matricula); err != nil {
			return d, err
		}

		return d, nil
	}

	if result.Next() {
		return d, nil
	}

	return domain.Dentista{}, errors.New("Registro não encontrado")
}

func (ds *dentistaStore) Update(id int, d domain.Dentista) (domain.Dentista, error) {

	_, err := ds.db.Exec("UPDATE dentistas SET nome=?, Sobrenome=?, matricula=? WHERE id=?", d.Nome, d.Sobrenome, d.Matricula, id)

	if err != nil {
		return d, err
	}

	return ds.GetById(id)
}

func (ds *dentistaStore) Delete(id int) error {

	result, err := ds.db.Exec("DELETE FROM dentistas WHERE id = ?", id)

	if err != nil {
		return err
	}

	count, err := result.RowsAffected()

	if err != nil {
		return errors.New("Dentista não cadastrado")
	}

	if count != 0 {
		return nil
	}

	return err
}
