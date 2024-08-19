package dbaser

import (
	"database/sql"
	"errors"
	"io/fs"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var createStatements = []string{
	`create table if not exists users (
id integer primary key,
email varchar(30) not null unique,
username varchar(20) not null unique,
password varchar(40) not null
);`,
	`create table if not exists user_profile (
id integer primary key,
user_id integer references users(id) on delete cascade,
content text
);`,
	`create table if not exists posts(
id integer primary key,
user_id integer references users(id),
content text not null
);`,
	`create table if not exists comments (
id integer primary key,
post_id integer references posts(id),
user_id integer references users(id),
content text not null
);`,
	`create table if not exists post_reaction (
id integer primary key,
user_id integer references users(id),
liked boolean
);`,
	`create table if not exists comment_reaction (
id integer primary key,
user_id integer references users(id),
liked boolean
);`,
	`create table if not exists categories (
id integer primary key,
label varchar(15) unique not null
);`,
}

func InitDb() {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			os.Create("forum.db")
			db, err = sql.Open("sqlite3", "./forum.db")
			if err != nil {
				log.Fatal("Error opening database file")
			} else {
				log.Fatal("Error opening database file")
			}
		}
	}
	defer db.Close()
	// Create tables.
	for _, statement := range createStatements {
		stmt, err := db.Prepare(statement)
		if err != nil {
			log.Fatal("Error preparing DB statement: ", statement)
		}
		if _, err = stmt.Exec(); err != nil {
			log.Fatal("Error creating database tables")
		}
	}
}
