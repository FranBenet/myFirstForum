package dbaser

import (
	"database/sql"
	"errors"
	"time"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(db *sql.DB, user models.User) (int, error) {
	if ok := UserEmailExists(db, user); !ok {
		return 0, errors.New("User e-mail already registered.")
	} else if ok = UsernameExists(db, user); !ok {
		return 0, errors.New("Username already registered.")
	}
	passHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 6)
	if err != nil {
		return 0, err
	}
	stmt, err := db.Prepare("insert into users (email, username, password) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.Email, user.Name, string(passHash))
	if err != nil {
		return 0, err
	}
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}

func UserEmailExists(db *sql.DB, user models.User) bool {
	row := db.QueryRow("select * from users where email=?", user.Email)
	if err := row.Scan(); err == sql.ErrNoRows {
		return false
	}
	return true
}

func UsernameExists(db *sql.DB, user models.User) bool {
	row := db.QueryRow("select * from users where username=?", user.Name)
	if err := row.Scan(); err == sql.ErrNoRows {
		return false
	}
	return true
}

func CheckPassword(db *sql.DB, user models.User) (bool, error) {
	var pass string
	row := db.QueryRow("select password from users where email=?", user.Email)
	if err := row.Scan(&pass); err == sql.ErrNoRows {
		return false, errors.New("User not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(user.Password)); err != nil {
		return false, err
	}
	return true, nil
}

func UserById(db *sql.DB, id int) (models.User, error) {
	var result models.User
	row := db.QueryRow("select id, email, username, created, avatar from users where id=?", id)
	var created string
	var avatar sql.NullString
	err := row.Scan(&result.Id, &result.Email, &result.Name, &created, &avatar)
	if err != nil {
		return models.User{}, err
	}
	timeCreated, err := time.Parse(time.RFC3339, created)
	if err != nil {
		return models.User{}, err
	}
	result.Created = timeCreated
	if avatar.Valid {
		result.Avatar = avatar.String
	}
	return result, nil
}
