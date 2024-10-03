package dbaser

import (
	"database/sql"
	"errors"
	"fmt"
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
password varchar(60) not null,
created datetime not null default (datetime('now', 'localtime')),
avatar varchar(30)
);`,
	`create table if not exists user_profile (
id integer primary key,
user_id integer references users(id) on delete cascade,
content text
);`,
	`create table if not exists posts(
id integer primary key,
user_id integer references users(id),
title varchar(60) not null,
content text not null,
created datetime not null default (datetime('now', 'localtime'))
);`,
	`create table if not exists comments (
id integer primary key,
post_id integer references posts(id),
user_id integer references users(id),
content text not null,
created datetime not null default (datetime('now', 'localtime'))
);`,
	`create table if not exists post_reactions (
post_id integer references posts(id),
user_id integer references users(id),
liked boolean,
primary key (post_id, user_id)
);`,
	`create table if not exists comment_reactions (
comment_id integer references comments(id),
user_id integer references users(id),
liked boolean,
primary key (comment_id, user_id)
);`,
	`create table if not exists categories (
id integer primary key,
label varchar(15) unique not null
);`,
	`create table if not exists post_categs (
post_id integer not null,
categ_id integer not null,
primary key (post_id, categ_id)
);`,
	`create table if not exists sessions (
id integer primary key,
user_id integer references users(id),
uuid varchar(40) not null,
expires datetime
);`,
}

func InitDb(path string) error {
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s", path))
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			log.Println("Creating database file:", path)
			os.Create(path)
			db, err = sql.Open("sqlite3", fmt.Sprintf("file:%s", path))
			if err != nil {
				return err
			}
			log.Println("Database file created!")
		}
	}
	defer db.Close()
	// Create tables.
	log.Println("Creating tables...")
	for _, statement := range createStatements {
		stmt, err := db.Prepare(statement)
		if err != nil {
			log.Println("Error preparing DB statement: ", statement)
			return err
		}
		defer stmt.Close()
		if _, err = stmt.Exec(); err != nil {
			log.Println("Error creating database table: ", err)
			return err
		}
	}
	log.Println("Tables created successfully!")
	return nil
}

func PopulateDb() {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal("Error opening database file")
	}
	defer db.Close()
	// Insert data.
	for table, statement := range insertStatements {
		stmt, err := db.Prepare(statement)
		if err != nil {
			log.Fatal("Error preparing DB statement: ", statement)
		}
		defer stmt.Close()
		res, err := stmt.Exec()
		if err != nil {
			log.Fatal("Error inserting data: ", err)
		}
		nrow, _ := res.RowsAffected()
		log.Printf("Number of rows inserted into table %s: %d\n", table, nrow)
	}
}
