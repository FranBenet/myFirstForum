package dbaser

import (
	"database/sql"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

func Categories(db *sql.DB) ([]models.Category, error) {
	var result []models.Category
	row, err := db.Query("select * from categories")
	if err == sql.ErrNoRows {
		return result, nil
	} else if err != nil {
		return []models.Category{}, err
	}
	for row.Next() {
		var cat models.Category
		err := row.Scan(&cat.Id, &cat.Name)
		if err != nil {
			return []models.Category{}, err
		}
		result = append(result, cat)
	}
	err = row.Err()
	if err != nil {
		return []models.Category{}, err
	}
	return result, nil
}

// Retrieve all categories associated with a specific post.
func PostCategories(db *sql.DB, id int) ([]models.Category, error) {
	var categories []models.Category
	row, err := db.Query("select categories.* from categories join post_categs on id=categ_id where post_id=?", id)
	if err == sql.ErrNoRows {
		return categories, nil
	} else if err != nil {
		return []models.Category{}, err
	}
	for row.Next() {
		var cat models.Category
		err := row.Scan(&cat.Id, &cat.Name)
		if err != nil {
			return []models.Category{}, err
		}
		categories = append(categories, cat)
	}
	err = row.Err()
	if err != nil {
		return []models.Category{}, err
	}
	return categories, nil
}

func AddPostCategory(db *sql.DB, category models.PostCategory) error {
	stmt, err := db.Prepare("insert or ignore into post_categs values (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(category.PostId, category.CategoryId)
	if err != nil {
		return err
	}
	return nil
}

func AddCategory(db *sql.DB, name string) (int, error) {
	stmt, err := db.Prepare("insert into categories (label) values (?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(name)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), nil
}

func CategoryExists(db *sql.DB, name string) (bool, error) {
	row := db.QueryRow("select label from categories where label=?", name)
	var label string
	if err := row.Scan(&label); err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func CategoryIdByName(db *sql.DB, name string) (int, error) {
	row := db.QueryRow("select id from categories where label=?", name)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func AddPostCategories(db *sql.DB, categories []string, postId int) error {
	for _, cat := range categories {
		exists, err := CategoryExists(db, cat)
		if err != nil {
			return err
		}
		if exists { // The category is present in the database.
			id, err := CategoryIdByName(db, cat)
			if err != nil {
				return err
			}
			err = AddPostCategory(db, models.PostCategory{PostId: postId, CategoryId: id})
			if err != nil {
				return err
			}
		} else { // New category added by the user.
			id, err := AddCategory(db, cat)
			if err != nil {
				return err
			}
			err = AddPostCategory(db, models.PostCategory{PostId: postId, CategoryId: id})
			if err != nil {
				return err
			}
		}
	}
	return nil
}
