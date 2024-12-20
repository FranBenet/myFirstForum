package dbaser

import (
	"database/sql"
	"log"
	"strings"
	"time"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

// All posts in the database.
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

func PostsByUser(db *sql.DB, userId int) ([]models.Post, error) {
	var result []models.Post
	row, err := db.Query("select * from posts where user_id=? order by created desc", userId)
	if err == sql.ErrNoRows {
		return result, nil
	} else if err != nil {
		return []models.Post{}, err
	}
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
func UserLikedPosts(db *sql.DB, userId int) ([]models.Post, error) {
	var result []models.Post
	row, err := db.Query("select posts.* from posts join post_reactions on posts.id=post_id join users on post_reactions.user_id=users.id where users.id=? and liked=? order by created desc", userId, 1)
	if err == sql.ErrNoRows {
		return result, nil
	} else if err != nil {
		return []models.Post{}, err
	}
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

func PostsByCategory(db *sql.DB, categoryId, page int) ([]models.Post, error) {
	var result []models.Post
	row, err := db.Query("select posts.* from posts join post_categs on post_id=posts.id where categ_id=? order by created desc", categoryId)
	if err == sql.ErrNoRows {
		return result, nil
	} else if err != nil {
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
	if err == sql.ErrNoRows {
		return result, nil
	} else if err != nil {
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

func PostById(db *sql.DB, postId int) (models.Post, error) {
	var result models.Post
	row := db.QueryRow("select * from posts where id=?", postId)
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

func NumberOfPosts(db *sql.DB) (int, error) {
	var num int
	row := db.QueryRow("select count(*) from posts;")
	if err := row.Scan(&num); err != nil {
		return 0, err
	}
	return num, nil
}

func Search(db *sql.DB, query string) ([]models.Post, error) {
	var result []models.Post
	if query == "" {
		return result, nil
	}
	posts, err := Posts(db)
	if err != nil {
		return result, err
	}
	for _, post := range posts {
		title := strings.ToLower(post.Title)
		content := strings.ToLower(post.Content)
		author, err := UserById(db, post.UserId)
		if err != nil {
			return result, err
		}
		username := strings.ToLower(author.Name)
		if strings.Contains(title, query) || strings.Contains(content, query) || strings.Contains(username, query) {
			result = append(result, post)
		}
	}
	return result, nil
}

func UntrendingPosts(db *sql.DB) ([]models.Post, error) {
	var result []models.Post
	row, err := db.Query("select posts.*, count(*) as num from posts join post_reactions on posts.id=post_id where liked=? group by post_id order by num desc", 0)
	if err == sql.ErrNoRows {
		return result, nil
	} else if err != nil {
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
