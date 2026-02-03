package configs

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB

func ConnectDB() {
	conn := os.Getenv("DATABASE_STRING")
	
	if conn == "" {
		log.Fatal("Data string not set")
	}

	db, err := sql.Open("sqlserver", conn)

	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	DB = db
}
