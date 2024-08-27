package dbaser

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

/* TODO
   - Filter posts by category. OK
   - Filter posts by reaction (likes).
   - Display all posts. OK
   - Filter posts by user? OK

Behaviour:
   When a post is requested from the DB, the idea is that the post will be displayed along with its tags, username,
   comments, likes/dislikes, etc. Should I create a function that will return all this? The other option is to have
   separate functions for each piece of information needed and I concatenate them for the handler function, which I
   guess is a more modular approach.
*/

func PostsByUser(user models.User) ([]models.Post, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal("Error opening database file")
	}
	defer db.Close()
	row, err := db.Query("select title, content, created from posts join users on user_id=users.id where users.email=?", user.Email)
	if err != nil {
		return []models.Post{}, err
	}
	var result []models.Post
	for row.Next() {
		var post models.Post
		var created string
		err = row.Scan(&post.Title, &post.Content, &created)
		if err != nil {
			return []models.Post{}, err
		}
		timeCreated, err := time.Parse(time.RFC3339, created)
		if err != nil {
			return []models.Post{}, err
		}
		post.Created = timeCreated.Format("02/01/2006 15:04")
		result = append(result, post)
	}
	return result, nil
}

func UserLikedPosts(user models.User) ([]models.Post, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal("Error opening database file")
	}
	defer db.Close()
	row, err := db.Query("select title, content, created from posts join post_reactions on posts.id=post_id join users on post_reactions.user_id=users.id where users.email=?", user.Email)
	if err != nil {
		return []models.Post{}, err
	}
	var result []models.Post
	for row.Next() {
		var post models.Post
		var created string
		err = row.Scan(&post.Title, &post.Content, &created)
		if err != nil {
			return []models.Post{}, err
		}
		timeCreated, err := time.Parse(time.RFC3339, created)
		if err != nil {
			return []models.Post{}, err
		}
		post.Created = timeCreated.Format("02/01/2006 15:04")
		result = append(result, post)
	}
	return result, nil
}

func Posts() ([]models.Post, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return []models.Post{}, err
	}
	defer db.Close()
	var result []models.Post
	row, err := db.Query("select title, content, created from posts order by created desc")
	if err != nil {
		return []models.Post{}, err
	}
	for row.Next() {
		var created string
		var post models.Post
		err := row.Scan(&post.Title, &post.Content, &created)
		timeCreated, err := time.Parse(time.RFC3339, created)
		if err != nil {
			log.Fatal("Error parsing post creation time")
		}
		post.Created = timeCreated.Format("02/01/2006 15:04")
		result = append(result, post)
	}
	return result, nil
}

func PostsByCategory(category models.Category) ([]models.Post, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return []models.Post{}, err
	}
	defer db.Close()
	var result []models.Post
	row, err := db.Query("select title, content, created from posts join post_categs on post_id=posts.id where categ_id=? order by created desc", category.Id)
	if err != nil {
		return []models.Post{}, err
	}
	for row.Next() {
		var created string
		var post models.Post
		err := row.Scan(&post.Title, &post.Content, &created)
		timeCreated, err := time.Parse(time.RFC3339, created)
		if err != nil {
			return []models.Post{}, err
		}
		post.Created = timeCreated.Format("02/01/2006 15:04")
		result = append(result, post)
	}
	return result, nil
}

func Categories() ([]models.Category, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return []models.Category{}, err
	}
	defer db.Close()
	var result []models.Category
	row, err := db.Query("select * from categories")
	if err != nil {
		return []models.Category{}, err
	}
	for row.Next() {
		var cat models.Category
		err := row.Scan(&cat.Id, &cat.Name)
		if err != nil {
			return []models.Category{}, err
		}
		result = append(result, cat)
	}
	return result, nil
}

func AddUser(user models.User) (int, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return 0, err
	}
	defer db.Close()
	passHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 6)
	if err != nil {
		return 0, err
	}
	stmt, err := db.Prepare("insert into users (email, username, password) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.Email, user.Name, string(passHash))
	if err != nil {
		return 0, err
	}
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}

func UserEmailExists(user models.User) (bool, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return true, err
	}
	defer db.Close()
	row := db.QueryRow("select * from users where email=?", user.Email)
	if err := row.Scan(); err == sql.ErrNoRows {
		return false, nil
	}
	return true, nil
}

func UsernameExists(user models.User) (bool, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return true, err
	}
	defer db.Close()
	row := db.QueryRow("select * from users where username=?", user.Name)
	if err := row.Scan(); err == sql.ErrNoRows {
		return false, nil
	}
	return true, nil
}

func CheckPassword(user models.User) (bool, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return false, err
	}
	defer db.Close()
	var pass string
	row := db.QueryRow("select password from users where email=?", user.Email)
	if err := row.Scan(&pass); err == sql.ErrNoRows {
		return false, errors.New("User not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(user.Password)); err != nil {
		return false, nil
	}
	return true, nil
}

func PostReactions(post models.Post) (int, int, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return 0, 0, err
	}
	defer db.Close()
	var likes, dislikes int
	row := db.QueryRow("select count(*) from post_reactions where post_id=? and liked=?", post.Id, 1)
	if err := row.Scan(&likes); err != nil {
		return 0, 0, err
	}
	row = db.QueryRow("select count(*) from post_reactions where post_id=? and liked=?", post.Id, 0)
	if err := row.Scan(&dislikes); err != nil {
		return 0, 0, err
	}
	return likes, dislikes, nil
}
