package dbaser

import (
	"database/sql"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

// TODO The user can either like or dislike a post, not both. So I have to check if there's already a reaction
// before inserting. If so, I update the entry.
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
func PostReactions(db *sql.DB, id int) (int, int, error) {
	var likes, dislikes int
	row := db.QueryRow("select count(*) from post_reactions where post_id=? and liked=?", id, 1)
	if err := row.Scan(&likes); err != nil {
		return 0, 0, err
	}
	row = db.QueryRow("select count(*) from post_reactions where post_id=? and liked=?", id, 0)
	if err := row.Scan(&dislikes); err != nil {
		return 0, 0, err
	}
	return likes, dislikes, nil
}

// TODO The user can either like or dislike a comment, not both. So I have to check if there's already a reaction
// before inserting. If so, I update the entry.
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

func PostLikeStatus(db *sql.DB, post_id, user_id int) (int, error) {
	row := db.QueryRow("select liked from post_reactions where post_id=? and user_id=?", post_id, user_id)
	var status int
	var liked bool
	if err := row.Scan(&liked); err == sql.ErrNoRows {
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	if liked {
		status = 1
	} else {
		status = -1
	}
	return status, nil
}
