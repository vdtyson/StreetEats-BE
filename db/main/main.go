package main

import (
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type PgConfig struct {
	user     string
	password string
	host     string
	port     string
	dbname   string
	sslmode  string
}

func main() {

}
