package dbaser

import (
	"database/sql"
	"fmt"
	"log"
)

func DbHandle(path string) (*sql.DB, error) {
	log.Println("Initialising database...")
	err := InitDb(path)
	if err != nil {
		return nil, err
	}
	//}
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s", path))
	return db, nil
}
