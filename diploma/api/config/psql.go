package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func GetPSQL() (*sql.DB, error) {
	host := os.Getenv("MONGO_HOST")
	port, _ := strconv.ParseInt(os.Getenv("PSQL_PORT"), 10, 64)
	user := os.Getenv("PSQL_USER")
	password := os.Getenv("PSQL_PASSWORD")
	dbName := os.Getenv("PSQL_DB_NAME")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected!")
	return db, nil
}
