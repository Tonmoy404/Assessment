package cmd

import (
	"database/sql"

	"github.com/Tonmoy404/Assessment/config"
)

func servreRest() {
	appConfig := config.GetApp()
	tableConfig := config.GetTable()

	db, err := sql.Open("postgres", "postgres://username:password@localhost/mydatabase?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

}
