package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

var Db *sqlx.DB

func InitDb() {
	dsn := os.Getenv("DSN")
	for i := 0; i < 5; i++ {
		db, err := sqlx.Connect("postgres", dsn)
		if err != nil {
			log.Println("Failed connecting to the database. Retrying in 2 seconds...")
			time.Sleep(time.Second * 2)
			continue
		} else {
			log.Println("Connected to the database")
			Db = db
			break
		}
	}
	if Db == nil {
		log.Fatal("Failed to connect to the database...")
	}

	Db.MustExec(schema)
}
