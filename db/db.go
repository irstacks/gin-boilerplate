package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

//DB ...
type DB struct {
	*sql.DB
}

const (
	//DbUser ...
	DbUser = "postgres"
	//DbPassword ...
	DbPassword = "postgres"
	//DbName ...
	DbName = "golang_gin_db"
)

var db *sql.DB

//Init ...
func Init() {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DbUser, DbPassword, DbName)

	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}

}

//ConnectDB ...
func ConnectDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

//GetDB ...
func GetDB() *sql.DB {
	return db
}
