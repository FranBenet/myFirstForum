package dbaser

import (
	"database/sql"
	"log"
	"time"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

/* TODO
   - Add a post. OK
   - Filter posts by category. OK
   - Filter posts by reaction (likes).
   - Display all posts. OK
   - Filter posts by user. OK
   - Number of comments of a post. OK
   - Trending posts (most likes).
   - User-liked posts. OK
   - Get categories associated with a post. OK
   - Get user ID from session ID.
   -

Behaviour:
   When a post is requested from the DB, the idea is that the post will be displayed along with its tags, username,
   comments, likes/dislikes, etc. Should I create a function that will return all this? The other option is to have
   separate functions for each piece of information needed and I concatenate them for the handler function, which I
   guess is a more modular approach.
*/

// Posts created by a specific user.
func PostsByUser(db *sql.DB, id int) ([]models.Post, error) {
	row, err := db.Query("select * from posts where user_id=?", id)
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
		post.Created = timeCreated
		result = append(result, post)
	}
	err = row.Err()
	if err != nil {
		return []models.Post{}, err
	}
	return result, nil
}

// All the posts liked by a specific user.
func UserLikedPosts(db *sql.DB, id int) ([]models.Post, error) {
	row, err := db.Query("select posts.* from posts join post_reactions on posts.id=post_id join users on post_reactions.user_id=users.id where users.id=? and liked=?", id, 1)
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
		post.Created = timeCreated
		result = append(result, post)
	}
	err = row.Err()
	if err != nil {
		return []models.Post{}, err
	}
	return result, nil
}

// All posts in the DB.
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
		post.Created = timeCreated
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
		post.Created = timeCreated
		result = append(result, post)
	}
	err = row.Err()
	if err != nil {
		return []models.Post{}, err
	}
	return result, nil
}

func AddPost(db *sql.DB, post models.Post) (int, error) {
	stmt, err := db.Prepare("insert into posts (title, content, user_id) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(post.Title, post.Content, post.UserId)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// Posts ranked by most likes.
func TrendingPosts(db *sql.DB, n int) ([]models.Post, error) {
	var result []models.Post
	row, err := db.Query("select posts.*, count(*) as num from posts join post_reactions on posts.id=post_id where liked=? group by post_id order by num desc limit ?", 1, n)
	if err != nil {
		return []models.Post{}, err
	}
	for row.Next() {
		var created string
		var post models.Post
		var count int
		err := row.Scan(&post.Id, &post.UserId, &post.Title, &post.Content, &created, &count)
		timeCreated, err := time.Parse(time.RFC3339, created)
		if err != nil {
			return []models.Post{}, err
		}
		post.Created = timeCreated
		result = append(result, post)
	}
	err = row.Err()
	if err != nil {
		return []models.Post{}, err
	}
	return result, nil
}

// Retrieve a specific post by its ID.
func PostById(db *sql.DB, id int) (models.Post, error) {
	var result models.Post
	row := db.QueryRow("select * from posts where id=?", id)
	var created string
	err := row.Scan(&result.Id, &result.UserId, &result.Title, &result.Content, &created)
	if err != nil {
		return models.Post{}, err
	}
	timeCreated, err := time.Parse(time.RFC3339, created)
	if err != nil {
		return models.Post{}, err
	}
	result.Created = timeCreated
	return result, nil
}
