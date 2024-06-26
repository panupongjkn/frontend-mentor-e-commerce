package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
	User       string
	Password   string
	Host       string
	Port       string
	DBName     string
	DisableTLS bool
}

func Init() *sql.DB {
	sslmode := "require"
	cfg := Config{
		User:       os.Getenv("DB_USERNAME"),
		Password:   os.Getenv("DB_PASSWORD"),
		Host:       os.Getenv("DB_DOMAIN"),
		Port:       os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DisableTLS: true,
	}
	if cfg.DisableTLS {
		sslmode = "disable"
	}
	var dataSoruce = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, sslmode)
	db, err := sql.Open("postgres", dataSoruce)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db
}
