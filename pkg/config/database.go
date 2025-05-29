package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitlializeDatabase() *sqlx.DB {
	host := GetString("DB_HOST", "localhost")
	port := GetInt("DB_PORT", 5432)
	user := GetString("DB_USER", "postgres")
	password := GetString("DB_PASSWORD", "admin")
	dbname := GetString("DB_NAME", "yummeals_app")
	sslmode := GetString("DB_SSL_MODE", "disable")

	var connString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	db, err := sqlx.Connect("postgres", connString)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
