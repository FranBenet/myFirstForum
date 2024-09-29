package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

func LikedPostsDataX(db *sql.DB, userId, page int) (models.MainPage, error) {
	var myPostsData models.MainPage

	pagination, err := NumberOfPagesX(db, userId)
	if err != nil {
		log.Println(err)
		myPostsData.Metadata.Error = err.Error()
		return myPostsData, err
	}
	log.Println("NUMBER OF PAGES: ", pagination)

	posts, err := dbaser.UserLikedPosts(db, userId, page)
	if err != nil {
		log.Print(err)
		myPostsData.Metadata.Error = err.Error()
		return myPostsData, err
	}
	log.Println("USERSPOSTS", posts)

	var postData []models.PostData
	for _, p := range posts {
		data, err := GetPostDataX(db, p, userId)
		if err != nil {
			log.Println(err)
			myPostsData.Metadata.Error = err.Error()
			return myPostsData, err
		}
		postData = append(postData, data)
	}

	if userId > 0 {
		myPostsData.Metadata.LoggedIn = true
		userData, err := dbaser.UserById(db, userId)
		if err != nil {
			log.Println(err)
			myPostsData.Metadata.Error = err.Error()
			return myPostsData, err
		}
		user := models.User{Avatar: userData.Avatar}
		myPostsData.User = user
	}

	pageData := models.Pagination{CurrentPage: page, TotalPages: pagination}
	myPostsData.Posts = postData
	myPostsData.Pagination = pageData

	return myPostsData, nil
}

func MyPostsDataX(db *sql.DB, userId, page int) (models.MainPage, error) {
	var myPostsData models.MainPage

	pagination, err := NumberOfPagesX(db, userId)
	if err != nil {
		log.Println(err)
		myPostsData.Metadata.Error = err.Error()
		return myPostsData, err
	}
	log.Println("NUMBER OF PAGES: ", pagination)

	posts, err := PostsByUserX(db, userId, page)
	if err != nil {
		log.Print(err)
		myPostsData.Metadata.Error = err.Error()
		return myPostsData, err
	}
	log.Println("USERSPOSTS", posts)

	var postData []models.PostData
	for _, p := range posts {
		data, err := GetPostDataX(db, p, userId)
		if err != nil {
			log.Println(err)
			myPostsData.Metadata.Error = err.Error()
			return myPostsData, err
		}
		postData = append(postData, data)
	}

	if userId > 0 {
		myPostsData.Metadata.LoggedIn = true
		userData, err := dbaser.UserById(db, userId)
		if err != nil {
			log.Println(err)
			myPostsData.Metadata.Error = err.Error()
			return myPostsData, err
		}
		user := models.User{Avatar: userData.Avatar}
		myPostsData.User = user
	}

	pageData := models.Pagination{CurrentPage: page, TotalPages: pagination}
	myPostsData.Posts = postData
	myPostsData.Pagination = pageData

	return myPostsData, nil
}

func NumberOfPagesX(db *sql.DB, userId int) (int, error) {
	nPosts, err := NumberOfPostsX(db, userId)
	if err != nil {
		log.Println("NPOSTS ERROR")
		return 0, err
	}

	var quot, rest int
	quot = nPosts / 5
	rest = nPosts % 5

	if quot == 0 {
		log.Println("1 PAGE")
		return 1, nil
	} else if quot > 0 && rest != 0 {
		log.Println("+1 PAGE")
		quot++
	}

	log.Println("NUMBER OF PAGES: ", quot)
	return quot, nil
}

func NumberOfPostsX(db *sql.DB, userId int) (int, error) {
	var num int
	row := db.QueryRow("select count(*) from post_reactions where user_id =?;", userId)
	if err := row.Scan(&num); err != nil {
		log.Println(err)
		return 0, err
	}

	log.Println("NUMBER OF POSTS: ", num)
	return num, nil
}

// Posts created by a specific user.
func PostsByUserX(db *sql.DB, userId int, page int) ([]models.Post, error) {
	log.Println("You are in the PostbyUserX")
	offset := (page - 1) * 5
	row, err := db.Query("select * from posts where user_id=? order by created desc limit 5 offset ?", userId, offset)
	if err != nil {
		log.Println("Error getting posts for the user")
		return []models.Post{}, err
	}
	var result []models.Post
	for row.Next() {
		var post models.Post
		var created string
		err = row.Scan(&post.Id, &post.UserId, &post.Title, &post.Content, &created)
		if err != nil {
			log.Println("Error scaning rows")
			return []models.Post{}, err
		}
		log.Println(post.Id, post.UserId, post.Title, post.Content, post.Created)
		timeCreated, err := time.Parse(time.RFC3339, created)
		if err != nil {
			log.Println("Error getting created time")
			return []models.Post{}, err
		}
		post.Created = timeCreated
		fmt.Println(post)
		result = append(result, post)
	}

	err = row.Err()
	if err != nil {
		log.Println("error error error")
		return []models.Post{}, err
	}
	return result, nil
}

func GetPostDataX(db *sql.DB, post models.Post, userId int) (models.PostData, error) {
	postUser, err := dbaser.UserById(db, post.UserId)
	if err != nil {
		return models.PostData{}, err
	}
	comments, err := dbaser.PostComments(db, post.Id)
	if err != nil {
		return models.PostData{}, err
	}
	likes, dislikes, err := dbaser.PostReactions(db, post.Id)
	if err != nil {
		return models.PostData{}, err
	}
	categories, err := dbaser.PostCategories(db, post.Id)
	if err != nil {
		return models.PostData{}, err
	}
	var likeStatus int
	if userId == 0 {
		likeStatus = 0
	} else {
		likeStatus, err = dbaser.PostLikeStatus(db, post.Id, userId)
		if err != nil {
			return models.PostData{}, err
		}
	}
	data := models.PostData{
		Post:         post,
		User:         postUser,
		Categories:   categories,
		Comments:     comments,
		LikeCount:    likes,
		DislikeCount: dislikes,
		Liked:        likeStatus,
	}
	return data, nil
}
