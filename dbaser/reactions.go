package dbaser

import (
	"database/sql"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

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
