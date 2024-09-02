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
   - Number of comments of a post. OK
   - Trending posts (number of likes and dislikes).
   - User-liked posts.
   - Get categories associated with a post.
   - Keep track of registered users.

Behaviour:
   When a post is requested from the DB, the idea is that the post will be displayed along with its tags, username,
   comments, likes/dislikes, etc. Should I create a function that will return all this? The other option is to have
   separate functions for each piece of information needed and I concatenate them for the handler function, which I
   guess is a more modular approach.
*/

func DbHandle(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Println("Error opening database file")
		return nil, err
	}
	return db, nil
}

func PostsByUser(db *sql.DB, user models.User) ([]models.Post, error) {
	row, err := db.Query("select posts.id, user_id, title, content, created from posts join users on user_id=users.id where users.email=?", user.Email)
	if err != nil {
		return []models.Post{}, err
	}
	var result []models.Post
	for row.Next() {
		var post models.Post
		var created string
		err = row.Scan(&post.Id, &post.UserId, &post.Title, &post.Content, &created)
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
	err = row.Err()
	if err != nil {
		return []models.Post{}, err
	}
	return result, nil
}

func UserLikedPosts(db *sql.DB, user models.User) ([]models.Post, error) {
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
	err = row.Err()
	if err != nil {
		return []models.Post{}, err
	}
	return result, nil
}

func Posts(db *sql.DB) ([]models.Post, error) {
	var result []models.Post
	row, err := db.Query("select * from posts order by created desc")
	if err != nil {
		return []models.Post{}, err
	}
	for row.Next() {
		var created string
		var post models.Post
		err := row.Scan(&post.Id, &post.UserId, &post.Title, &post.Content, &created)
		timeCreated, err := time.Parse(time.RFC3339, created)
		if err != nil {
			log.Fatal("Error parsing post creation time")
		}
		post.Created = timeCreated.Format("02/01/2006 15:04")
		result = append(result, post)
	}
	err = row.Err()
	if err != nil {
		return []models.Post{}, err
	}
	return result, nil
}

func PostsByCategory(db *sql.DB, category models.Category) ([]models.Post, error) {
	var result []models.Post
	row, err := db.Query("select * from posts join post_categs on post_id=posts.id where categ_id=? order by created desc", category.Id)
	if err != nil {
		return []models.Post{}, err
	}
	for row.Next() {
		var created string
		var post models.Post
		err := row.Scan(&post.Id, &post.UserId, &post.Title, &post.Content, &created)
		timeCreated, err := time.Parse(time.RFC3339, created)
		if err != nil {
			return []models.Post{}, err
		}
		post.Created = timeCreated.Format("02/01/2006 15:04")
		result = append(result, post)
	}
	err = row.Err()
	if err != nil {
		return []models.Post{}, err
	}
	return result, nil
}

func Categories(db *sql.DB) ([]models.Category, error) {
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
	err = row.Err()
	if err != nil {
		return []models.Category{}, err
	}
	return result, nil
}

func AddUser(db *sql.DB, user models.User) (int, error) {
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

func UserEmailExists(db *sql.DB, user models.User) (bool, error) {
	row := db.QueryRow("select * from users where email=?", user.Email)
	if err := row.Scan(); err == sql.ErrNoRows {
		return false, nil
	}
	return true, nil
}

func UsernameExists(db *sql.DB, user models.User) (bool, error) {
	row := db.QueryRow("select * from users where username=?", user.Name)
	if err := row.Scan(); err == sql.ErrNoRows {
		return false, nil
	}
	return true, nil
}

func CheckPassword(db *sql.DB, user models.User) (bool, error) {
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

func PostReactions(db *sql.DB, post models.Post) (int, int, error) {
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

func CommentNumber(db *sql.DB, post models.Post) (int, error) {
	var result int
	row := db.QueryRow("select count(*) from comments where post_id=?", post.Id)
	if err := row.Scan(&result); err != nil {
		return 0, err
	}
	return result, nil
}
