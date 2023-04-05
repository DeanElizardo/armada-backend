package databaseServices

import (
	"armadabackend/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s "+
	"dbname=%s sslmode=disable",
	config.Host, config.Port, config.User, config.Password, config.DBname)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		e := fmt.Errorf("ConnectDB: open: %s", err)
		return db, e
	}
	//db.Close() //Should probably call this on the returned DB in the query funcs
	err = db.Ping()
	if err != nil {
		e := fmt.Errorf("ConnectDB: ping: %s", err)
		return db, e
	}

	return db, nil
}
