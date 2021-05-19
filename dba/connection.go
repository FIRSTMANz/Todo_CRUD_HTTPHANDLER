package dba

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	dbConn *sql.DB
)

func CloseConnection() {
	if dbConn != nil {
		dbConn.Close()
	}
}

func NewConnection() (*sql.DB, error) {
	if dbConn == nil {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			return nil, err
		}
		dbConn = db
	}

	return dbConn, nil
}
