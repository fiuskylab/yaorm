package main

import (
	"github.com/fiuskylab/yaorm"
	"github.com/fiuskylab/yaorm/driver"
	"github.com/fiuskylab/yaorm/driver/pgsql"
)

func main() {
	db, err := yaorm.Open(driver.PGSQL, &pgsql.DSN{
		User:     "user",
		Password: "password",
		Host:     "localhost",
		Port:     5432,
		DBName:   "db_name",
		Params: map[string]string{
			"sslmode": "disable",
		},
	})

	if err != nil {
		// handle err
		panic(err)
	}

	if err := db.Ping(); err != nil {
		// implemente reconnection logic
		panic(err)
	}
}
