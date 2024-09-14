package dbaser

import (
	"database/sql"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

func Categories(db *sql.DB) ([]models.Category, error) {
	var result []models.Category
	row, err := db.Query("select * from categories")
	if err != nil {
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
	if err != nil {
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