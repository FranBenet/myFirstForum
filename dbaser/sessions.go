package dbaser

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"time"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

func AddSession(db *sql.DB, user models.User) (string, error) {
	stmt, err := db.Prepare("insert into sessions (user_id, uuid, expires) values (?, ?, ?)")
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	uuid, err := GenerateUuid(16)
	if err != nil {
		return "", err
	}
	expiration := time.Now().Add(time.Hour).Format("2006-01-02 15:04:05")
	_, err = stmt.Exec(user.Id, uuid, expiration)
	if err != nil {
		return "", err
	}
	return uuid, nil
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

func DeleteSession(db *sql.DB, uuid string) (int, error) {
	row, err := db.Exec("delete from sessions where uuid=?", uuid)
	if err != nil {
		return 0, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
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

func ValidSession(db *sql.DB, uuid string) (bool, error) {
	if uuid == "" {
		return false, nil
	}
	row := db.QueryRow("select expires from sessions where uuid=?", uuid)
	var expires string
	if err := row.Scan(&expires); err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}
	deadline, err := time.Parse(time.RFC3339, expires)
	if err != nil {
		return false, err
	}
	if deadline.Before(time.Now()) {
		return false, nil
	}
	return true, nil
}

func GenerateUuid(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(buffer)[:length], nil
}
