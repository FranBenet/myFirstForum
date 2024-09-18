package dbaser

import (
	"database/sql"
	"fmt"
	"log"
)

func DbHandle(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s", path))
	if err != nil {
		log.Println("Error opening database file")
		return nil, err
	}
	return db, nil
}
