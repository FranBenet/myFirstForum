package dbaser

import (
	"database/sql"
	"errors"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

// Adds a post reaction and returns the ID of the inserted row.
func AddPostReaction(db *sql.DB, reaction models.PostReaction) (int, error) {
	exists, err := PostReactionExists(db, reaction)
	if err != nil {
		return 0, err
	} else if exists {
		currentReaction, err := GetPostReaction(db, reaction)
		if err != nil {
			return 0, err
		}
		if currentReaction.Liked == reaction.Liked {
			affectedRows, err := DeletePostReaction(db, reaction)
			if err != nil {
				return 0, err
			} else {
				return affectedRows, nil
			}
		} else {
			affectedRows, err := UpdatePostReaction(db, reaction)
			if err != nil {
				return 0, err
			} else {
				return affectedRows, nil
			}
		}
	}
	stmt, err := db.Prepare("insert into post_reactions values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
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

func DeletePostReaction(db *sql.DB, reaction models.PostReaction) (int, error) {
	exists, err := PostReactionExists(db, reaction)
	if err != nil {
		return 0, err
	} else if !exists {
		return 0, errors.New("Post reaction not found.")
	}
	stmt, err := db.Prepare("delete from post_reactions where post_id=? and user_id=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(reaction.PostId, reaction.UserId)
	if err != nil {
		return 0, err
	}
	nRows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(nRows), nil
}

func UpdatePostReaction(db *sql.DB, reaction models.PostReaction) (int, error) {
	exists, err := PostReactionExists(db, reaction)
	if err != nil {
		return 0, err
	} else if !exists {
		return 0, errors.New("Post reaction not found.")
	}
	stmt, err := db.Prepare("update post_reactions set liked=? where post_id=? and user_id=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(reaction.Liked, reaction.PostId, reaction.UserId)
	if err != nil {
		return 0, err
	}
	nRows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(nRows), nil
}

// Returns the number of likes and dislikes of a post.
func PostReactions(db *sql.DB, id int) (int, int, error) {
	var likes, dislikes int
	row := db.QueryRow("select count(*) from post_reactions where post_id=? and liked=?", id, true)
	if err := row.Scan(&likes); err != nil {
		return 0, 0, err
	}
	row = db.QueryRow("select count(*) from post_reactions where post_id=? and liked=?", id, false)
	if err := row.Scan(&dislikes); err != nil {
		return 0, 0, err
	}
	return likes, dislikes, nil
}

func AddCommentReaction(db *sql.DB, reaction models.CommentReaction) (int, error) {
	exists, err := CommentReactionExists(db, reaction)
	if err != nil {
		return 0, err
	} else if exists {
		currentReaction, err := GetCommentReaction(db, reaction)
		if err != nil {
			return 0, err
		}
		if currentReaction.Liked == reaction.Liked {
			DeleteCommentReaction(db, reaction)
		} else {
			UpdateCommentReaction(db, reaction)
		}
	}

	stmt, err := db.Prepare("insert into comment_reactions values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
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

func DeleteCommentReaction(db *sql.DB, reaction models.CommentReaction) (int, error) {
	exists, err := CommentReactionExists(db, reaction)
	if err != nil {
		return 0, err
	} else if !exists {
		return 0, errors.New("Comment reaction not found.")
	}
	stmt, err := db.Prepare("delete from commentt_reactions where commentt_id=? and user_id=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(reaction.CommentId, reaction.UserId)
	if err != nil {
		return 0, err
	}
	nRows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(nRows), nil
}

func UpdateCommentReaction(db *sql.DB, reaction models.CommentReaction) (int, error) {
	exists, err := CommentReactionExists(db, reaction)
	if err != nil {
		return 0, err
	} else if !exists {
		return 0, errors.New("Comment reaction not found.")
	}
	stmt, err := db.Prepare("update comment_reactions set liked=? where post_id=? and user_id=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(reaction.Liked, reaction.CommentId, reaction.UserId)
	if err != nil {
		return 0, err
	}
	nRows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(nRows), nil
}

// Returns the number of likes and dislikes of a comment.
func CommentReactions(db *sql.DB, id int) (int, int, error) {
	var likes, dislikes int
	row := db.QueryRow("select count(*) from comment_reactions where comment_id=? and liked=?", id, 1)
	if err := row.Scan(&likes); err != nil {
		return 0, 0, err
	}
	row = db.QueryRow("select count(*) from comment_reactions where comment_id=? and liked=?", id, 0)
	if err := row.Scan(&dislikes); err != nil {
		return 0, 0, err
	}
	return likes, dislikes, nil
}

// Determine if a user has either liked or disliked a post.
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

// Determine if a user has either liked or disliked a comment.
func CommentLikeStatus(db *sql.DB, comment_id, user_id int) (int, error) {
	row := db.QueryRow("select liked from comment_reactions where comment_id=? and user_id=?", comment_id, user_id)
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

func PostReactionExists(db *sql.DB, reaction models.PostReaction) (bool, error) {
	var count int
	row := db.QueryRow("select count(*) from post_reactions where post_id=? and user_id=?", reaction.PostId, reaction.UserId)
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func GetPostReaction(db *sql.DB, reaction models.PostReaction) (models.PostReaction, error) {
	var result models.PostReaction
	row := db.QueryRow("select * from post_reactions where post_id=? and user_id=?", reaction.PostId, reaction.UserId)
	if err := row.Scan(&result.PostId, &result.UserId, &result.Liked); err != nil {
		return models.PostReaction{}, err
	}
	return result, nil
}

func CommentReactionExists(db *sql.DB, reaction models.CommentReaction) (bool, error) {
	var count int
	row := db.QueryRow("select count(*) from comment_reactions where comment_id=? and user_id=?", reaction.CommentId, reaction.UserId)
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func GetCommentReaction(db *sql.DB, reaction models.CommentReaction) (models.CommentReaction, error) {
	var result models.CommentReaction
	row := db.QueryRow("select * from commentt_reactions where comment_id=? and user_id=?", reaction.CommentId, reaction.UserId)
	if err := row.Scan(&result.CommentId, &result.UserId, &result.Liked); err != nil {
		return models.CommentReaction{}, err
	}
	return result, nil
}
