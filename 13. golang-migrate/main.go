package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB_URL = "postgresql://postgres:example@localhost:5432/golang-migrate?sslmode=disable"
var MIGRATION_FILE = "file:///home/istiyak/Programs/Github/go-practices/13. golang-migrate/migrations"

func fromDatabaseURL() {
	m, err := migrate.New(
		MIGRATION_FILE,
		DB_URL,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Make migration up
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Migration Up finished")

	// Make migration down
	// if err := m.Down(); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Migration Down finished")
}


func fromExistingDB(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		MIGRATION_FILE,
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	// Make migration Up
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Migration Up finished")

	// Make migration down
	// if err := m.Down(); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Migration Down finished")
}

func main() {
	// Migrate up using existing database connection
	// db, err := sql.Open("postgres", DB_URL)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fromExistingDB(db) 


	// Migrate from the database connection url
	fromDatabaseURL()


}




