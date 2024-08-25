package dbaser

import (
	"database/sql"
	"errors"
	"io/fs"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

/* TODO
   - Filter posts by category.
   - Filter posts by reaction (likes).
   - Display all posts.
   - Filter posts by user?

Behaviour:
   When a post is requested from the DB, the idea is that the post will be displayed along with its tags, username,
   comments, likes/dislikes, etc. Should I create a function that will return all this? The other option is to have
   separate functions for each piece of information needed and I concatenate them for the handler function, which I
   guess is a more modular approach.
*/

var createStatements = []string{
	`create table if not exists users (
id integer primary key,
email varchar(30) not null unique,
username varchar(20) not null unique,
password varchar(30) not null
);`,
	`create table if not exists user_profile (
id integer primary key,
user_id integer references users(id) on delete cascade,
content text
);`,
	`create table if not exists posts(
id integer primary key,
user_id integer references users(id),
created datetime not null default current_timestamp,
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
	`create table if not exists post_categs (
categ_id integer not null,
post_id integer not null,
primary key (categ_id, post_id)
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

func PopulateDb() {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal("Error opening database file")
	}
	defer db.Close()
	// Insert users.
	stmt, err := db.Prepare(insertUsers)
	if err != nil {
		log.Fatal("Error preparing (user) insert statement: ", stmt)
	}
	defer stmt.Close()
	res, err := stmt.Exec()
	if err != nil {
		log.Fatal("Error inserting into database: ", err)
	}
	nrow, _ := res.RowsAffected()
	log.Println("Number of rows inserted into table users: ", nrow)
	// Insert posts.
	stmt, err = db.Prepare(insertPosts)
	if err != nil {
		log.Fatal("Error preparing (posts) insert statement: ", stmt)
	}
	res, err = stmt.Exec()
	if err != nil {
		log.Fatal("Error inserting into database: ", err)
	}
	nrow, _ = res.RowsAffected()
	log.Println("Number of rows inserted into table posts: ", nrow)
	// Insert categories.
	stmt, err = db.Prepare(insertCategs)
	if err != nil {
		log.Fatal("Error preparing (categories) insert statement: ", stmt)
	}
	res, err = stmt.Exec()
	if err != nil {
		log.Fatal("Error inserting into database: ", err)
	}
	nrow, _ = res.RowsAffected()
	log.Println("Number of rows inserted into table categories: ", nrow)
	// Insert post categories.
	stmt, err = db.Prepare(postCategories)
	if err != nil {
		log.Fatal("Error preparing (post categories) insert statement: ", stmt)
	}
	res, err = stmt.Exec()
	if err != nil {
		log.Fatal("Error inserting into database: ", err)
	}
	nrow, _ = res.RowsAffected()
	log.Println("Number of rows inserted into table post_categs: ", nrow)
}
