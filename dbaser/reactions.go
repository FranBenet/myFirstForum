package dbaser

import (
	"database/sql"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

func AddPostReaction(db *sql.DB, reaction models.PostReaction) (int, error) {
	stmt, err := db.Prepare("insert into post_reactions values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(reaction.PostId, reaction.UserId, reaction.Liked)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// Returns the number of likes and dislikes of a post.
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

func AddCommentReaction(db *sql.DB, reaction models.CommentReaction) (int, error) {
	stmt, err := db.Prepare("insert into comment_reactions values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(reaction.CommentId, reaction.UserId, reaction.Liked)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// Returns the number of likes and dislikes of a comment.
func CommentReactions(db *sql.DB, comment models.Comment) (int, int, error) {
	var likes, dislikes int
	row := db.QueryRow("select count(*) from comment_reactions where comment_id=? and liked=?", comment.Id, 1)
	if err := row.Scan(&likes); err != nil {
		return 0, 0, err
	}
	row = db.QueryRow("select count(*) from comment_reactions where comment_id=? and liked=?", comment.Id, 0)
	if err := row.Scan(&dislikes); err != nil {
		return 0, 0, err
	}
	return likes, dislikes, nil
}
