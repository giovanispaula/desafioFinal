package store

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Gio*2802@/desafio_backend")

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")

	return db, nil
}
