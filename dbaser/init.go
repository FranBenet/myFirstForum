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
password text not null
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
created datetime not null default current_timestamp
);`,
	`create table if not exists comments (
id integer primary key,
post_id integer references posts(id),
user_id integer references users(id),
content text not null,
created datetime not null default current_timestamp
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
		defer stmt.Close()
		if _, err = stmt.Exec(); err != nil {
			log.Fatal("Error creating database table: ", err)
		}
	}
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
