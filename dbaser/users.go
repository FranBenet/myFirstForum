package dbaser

import (
	"database/sql"
	"errors"
	"time"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(db *sql.DB, user models.User) (int, error) {
	if exists, err := UserEmailExists(db, user.Email); exists {
		return 0, errors.New("User e-mail already registered.")
	} else if err != nil {
		return 0, err
	} else if exists, err = UsernameExists(db, user.Name); exists {
		return 0, errors.New("Username already registered.")
	} else if err != nil {
		return 0, err
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

func UserEmailExists(db *sql.DB, email string) (bool, error) {
	row := db.QueryRow("select email from users where email=?", email)
	var epost string
	if err := row.Scan(&epost); err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func UsernameExists(db *sql.DB, username string) (bool, error) {
	row := db.QueryRow("select username from users where username=?", username)
	var name string
	if err := row.Scan(&name); err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func CheckPassword(db *sql.DB, email, password string) (bool, error) {
	var pass string
	row := db.QueryRow("select password from users where email=?", email)
	if err := row.Scan(&pass); err == sql.ErrNoRows {
		return false, errors.New("User not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(password)); err != nil {
		return false, errors.New("Incorrect password")
	}
	return true, nil
}

func UserById(db *sql.DB, id int) (models.User, error) {
	var result models.User
	row := db.QueryRow("select id, email, username, created, avatar from users where id=?", id)
	var created string
	var avatar sql.NullString
	if err := row.Scan(&result.Id, &result.Email, &result.Name, &created, &avatar); err != nil {
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

func UserByEmail(db *sql.DB, email string) (models.User, error) {
	var result models.User
	row := db.QueryRow("select id, email, username, created, avatar from users where email=?", email)
	var created string
	var avatar sql.NullString
	if err := row.Scan(&result.Id, &result.Email, &result.Name, &created, &avatar); err != nil {
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

func ValidateLogin(db *sql.DB, email, password string) (models.User, error) {
	if emailOk, err := UserEmailExists(db, email); !emailOk {
		return models.User{}, errors.New("User e-mail not found!")
	} else if err != nil {
		return models.User{}, err
	}
	if passOk, err := CheckPassword(db, email, password); !passOk {
		return models.User{}, errors.New("Incorrect password!")
	} else if err != nil {
		return models.User{}, err
	}
	user, err := UserByEmail(db, email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
