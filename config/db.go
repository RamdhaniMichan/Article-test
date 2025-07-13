package config

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres",
		"host="+os.Getenv("DB_HOST")+
			" user="+os.Getenv("DB_USER")+
			" password="+os.Getenv("DB_PASS")+
			" dbname="+os.Getenv("DB_NAME")+
			" sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	maxOpen, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	if maxOpen == 0 {
		maxOpen = 25
	}
	maxIdle, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	if maxIdle == 0 {
		maxIdle = 25
	}
	maxLifetime, _ := strconv.Atoi(os.Getenv("DB_CONN_MAX_LIFETIME_MIN"))
	if maxLifetime == 0 {
		maxLifetime = 5
	}

	db.SetMaxOpenConns(maxOpen)
	db.SetMaxIdleConns(maxIdle)
	db.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Minute)

	return db
}
