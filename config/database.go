package config

import (
	"assignTele/helper"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "blogs"
)

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", sqlInfo)
	err = db.Ping()
	helper.PanicIfError(err)

	log.Info().Msg("Ping Successful, Connected to Database!")

	return db
}
