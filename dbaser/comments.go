package dbaser

import (
	"database/sql"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

func AddComment(db *sql.DB, comment models.Comment, post models.Post, user models.User) (int, error) {
	stmt, err := db.Prepare("insert into comments (post_id, user_id, content) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(post.Id, user.Id, comment.Content)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func CommentNumber(db *sql.DB, post models.Post) (int, error) {
	var result int
	row := db.QueryRow("select count(*) from comments where post_id=?", post.Id)
	if err := row.Scan(&result); err != nil {
		return 0, err
	}
	return result, nil
}
