package dbaser

import (
	"database/sql"
	"time"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

/* TODO
   - Add session. OK
   - Delete session by UUID.
   - Get user by session UUID. OK
*/

func AddSession(db *sql.DB, session models.Session) (int, error) {
	stmt, err := db.Prepare("insert into sessions (user_id, uuid, expires) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	expiration := session.ExpiresAt.Format("2006-01-02 15:04:05")
	res, err := stmt.Exec(session.UserId, session.Uuid, expiration)
	if err != nil {
		return 0, err
	}
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}

func SessionById(db *sql.DB, id int) (models.Session, error) {
	var result models.Session
	row := db.QueryRow("select * from sessions where id=?", id)
	var expiration string
	if err := row.Scan(&result.Id, &result.UserId, &result.Uuid, &expiration); err != nil {
		return models.Session{}, err
	}
	timeCreated, err := time.Parse(time.RFC3339, expiration)
	if err != nil {
		return models.Session{}, err
	}
	result.ExpiresAt = timeCreated
	return result, nil
}

func DeleteSessionUuid(db *sql.DB, uuid string) error {
	_, err := db.Exec("delete from sessions where uuid=?", uuid)
	if err != nil {
		return err
	}
	return nil
}

func SessionUser(db *sql.DB, uuid string) (int, error) {
	if uuid == "" {
		return 0, nil
	}
	row := db.QueryRow("select user_id from sessions where uuid=?", uuid)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
