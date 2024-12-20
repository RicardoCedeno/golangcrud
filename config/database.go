package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDatabase() (*sql.DB, error) {
	connString := "user=postgres password=7dejulio dbname=golang sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("error al conectar a la base de datos ", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error al conectar a la base de datos ", err)
	}

	fmt.Println("Conexión exitosa a la base de datos")

	return db, nil
}
