package dbaser

import (
	"database/sql"
	"log"
	"time"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

func PostsByUser(email string) (models.Post, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal("Error opening database file")
	}
	defer db.Close()
	var post models.Post
	var created string
	row := db.QueryRow("select content, created from posts join users on user_id=users.id where users.email=?", email)
	if err := row.Scan(&post.Content, &created); err != nil {
		return models.Post{}, err
	}
	timeCreated, err := time.Parse(time.RFC3339, created)
	if err != nil {
		return models.Post{}, err
	}
	post.Created = timeCreated.Format("02/01/2006 15:04")
	return post, nil
}

func AllPosts() ([]models.Post, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return []models.Post{}, err
	}
	defer db.Close()
	var result []models.Post
	row, err := db.Query("select content, created from posts order by created desc")
	if err != nil {
		return []models.Post{}, err
	}
	for row.Next() {
		var created string
		var post models.Post
		err := row.Scan(&post.Content, &created)
		timeCreated, err := time.Parse(time.RFC3339, created)
		if err != nil {
			log.Fatal("Error parsing post creation time")
		}
		post.Created = timeCreated.Format("02/01/2006 15:04")
		result = append(result, post)
	}
	return result, nil
}
