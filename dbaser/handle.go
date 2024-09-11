package dbaser

import (
	"database/sql"
	"log"
)

func DbHandle(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Println("Error opening database file")
		return nil, err
	}
	return db, nil
}
