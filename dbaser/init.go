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

type User struct {
	Email    string
	Name     string
	Password string
}

type Post struct {
	UserId  int
	Created string
	Content string
}

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
created datetime,
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

func PopulateDb() {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal("Error opening database file")
	}
	defer db.Close()
	for _, user := range users {
		stmt, err := db.Prepare("insert into users(email, username, password) values (?, ?, ?)")
		if err != nil {
			log.Fatal("Error preparing insert statement: ", stmt)
		}
		defer stmt.Close()
		res, err := stmt.Exec(user.Email, user.Name, user.Password)
		if err != nil {
			log.Fatal("Error inserting into database: ", err)
		}
		id, _ := res.LastInsertId()
		log.Println("Successfully inserted into table users: ", id, user.Email, user.Name)
	}
	for _, post := range posts {
		stmt, err := db.Prepare("insert into posts(user_id, created, content) values (?, ?, ?)")
		if err != nil {
			log.Fatal("Error preparing insert statement: ", stmt)
		}
		defer stmt.Close()
		res, err := stmt.Exec(post.UserId, post.Created, post.Content)
		if err != nil {
			log.Fatal("Error inserting into database: ", err)
		}
		id, _ := res.LastInsertId()
		log.Println("Successfully inserted into table posts: ", id, post.Created, post.Content)
	}
}

func GetUser(email string) User {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal("Error opening database file")
	}
	defer db.Close()
	var usr User
	row := db.QueryRow("select email, username from users where username=?", email)
	if err := row.Scan(&usr.Email, &usr.Name); err != nil {
		log.Fatal("Error querying database: ", err)
	}
	return usr
}
