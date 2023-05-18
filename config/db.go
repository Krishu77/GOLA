package config

import (
	"log"
	"os"

	controllers "GOLA/controllers"

	"github.com/go-pg/pg/v9"
)

// Connecting to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "bujjigadu",
		Addr:     "localhost:5432",
		Database: "testdb",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	//controllers.CreateTodoTable(db)
	controllers.InitiateDB(db)
	return db
}
