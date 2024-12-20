package dbaser

import (
	"database/sql"
	"time"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

func AddComment(db *sql.DB, comment models.Comment) (int, error) {
	stmt, err := db.Prepare("insert into comments (post_id, user_id, content) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(comment.PostId, comment.UserId, comment.Content)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func PostComments(db *sql.DB, id int) ([]models.Comment, error) {
	var result []models.Comment
	row, err := db.Query("select * from comments where post_id=? order by created", id)
	if err == sql.ErrNoRows {
		return result, nil
	} else if err != nil {
		return []models.Comment{}, err
	}
	for row.Next() {
		var comment models.Comment
		var created string
		err = row.Scan(&comment.Id, &comment.PostId, &comment.UserId, &comment.Content, &created)
		if err != nil {
			return []models.Comment{}, err
		}
		timeCreated, err := time.Parse(time.RFC3339, created)
		if err != nil {
			return []models.Comment{}, err
		}
		comment.Created = timeCreated
		result = append(result, comment)
	}
	err = row.Err()
	if err != nil {
		return []models.Comment{}, err
	}
	return result, nil
}

func CommentNumber(db *sql.DB, post models.Post) (int, error) {
	var result int
	row := db.QueryRow("select count(*) from comments where post_id=?", post.Id)
	if err := row.Scan(&result); err != nil {
		return 0, err
	}
	return result, nil
}

func CommentById(db *sql.DB, id int) (models.Comment, error) {
	var result models.Comment
	row := db.QueryRow("select * from comments where id=?", id)
	var created string
	err := row.Scan(&result.Id, &result.PostId, &result.UserId, &result.Content, &created)
	if err != nil {
		return models.Comment{}, err
	}
	timeCreated, err := time.Parse(time.RFC3339, created)
	if err != nil {
		return models.Comment{}, err
	}
	result.Created = timeCreated
	return result, nil
}
