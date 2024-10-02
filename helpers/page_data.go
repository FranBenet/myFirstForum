package helpers

import (
	"database/sql"
	"log"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

func MainPageData(db *sql.DB, userId, page int) (models.MainPage, error) {
	var mainData models.MainPage
	posts, err := dbaser.Posts(db)
	if err != nil {
		log.Print(err)
		mainData.Metadata.Error = err.Error()
		return mainData, err
	}
	mainData, err = PageData(db, posts, "main", userId, page)
	if err != nil {
		return mainData, err
	}
	return mainData, nil
}

// Similar to MainPageData but relative to a single post.
func PostPageData(db *sql.DB, postId, sessionUser int) (models.PostPage, error) {
	var postData models.PostPage
	post, err := dbaser.PostById(db, postId)
	if err != nil {
		postData.Metadata.Error = err.Error()
		return postData, err
	}
	data, err := GetPostData(db, post, sessionUser)
	if err != nil {
		postData.Metadata.Error = err.Error()
		return postData, err
	}
	var comments []models.CommentData
	for _, comment := range data.Comments {
		commData, err := GetCommentData(db, comment, sessionUser)
		if err != nil {
			postData.Metadata.Error = err.Error()
			return postData, err
		}
		comments = append(comments, commData)
	}
	if sessionUser > 0 {
		postData.Metadata.LoggedIn = true
		userData, err := dbaser.UserById(db, sessionUser)
		if err != nil {
			postData.Metadata.Error = err.Error()
			return postData, err
		}
		user := models.User{Avatar: userData.Avatar}
		postData.User = user
	}
	postData.Post = data
	postData.Comments = comments
	return postData, nil
}

func MyPostsPageData(db *sql.DB, userId, page int) (models.MainPage, error) {
	var mainData models.MainPage
	posts, err := dbaser.PostsByUser(db, userId)
	if err != nil {
		log.Print(err)
		mainData.Metadata.Error = err.Error()
		return mainData, err
	}
	mainData, err = PageData(db, posts, "", userId, page)
	if err != nil {
		return mainData, err
	}
	return mainData, nil
}

func MyLikedPostsPageData(db *sql.DB, userId, page int) (models.MainPage, error) {
	var mainData models.MainPage
	posts, err := dbaser.UserLikedPosts(db, userId)
	if err != nil {
		log.Print(err)
		mainData.Metadata.Error = err.Error()
		return mainData, err
	}
	mainData, err = PageData(db, posts, "", userId, page)
	if err != nil {
		return mainData, err
	}
	return mainData, nil
}

func SearchPageData(db *sql.DB, query string, userId, page int) (models.MainPage, error) {
	var mainData models.MainPage
	posts, err := dbaser.Search(db, query)
	if err != nil {
		log.Print(err)
		mainData.Metadata.Error = err.Error()
		return mainData, err
	}
	mainData, err = PageData(db, posts, "filter", userId, page)
	if err != nil {
		return mainData, err
	}
	return mainData, nil
}

func TrendingPageData(db *sql.DB, userId, page int) (models.MainPage, error) {
	var mainData models.MainPage
	posts, err := dbaser.TrendingPosts(db, -1)
	if err != nil {
		log.Print(err)
		mainData.Metadata.Error = err.Error()
		return mainData, err
	}
	mainData, err = PageData(db, posts, "main", userId, page)
	if err != nil {
		return mainData, err
	}
	return mainData, nil
}

func UntrendingPageData(db *sql.DB, userId, page int) (models.MainPage, error) {
	var mainData models.MainPage
	posts, err := dbaser.UntrendingPosts(db)
	if err != nil {
		log.Print(err)
		mainData.Metadata.Error = err.Error()
		return mainData, err
	}
	mainData, err = PageData(db, posts, "main", userId, page)
	if err != nil {
		return mainData, err
	}
	return mainData, nil
}

// Styles that require trending: main, filters, search; NO my posts, liked posts.
// Styles that require categories: main, category filter result.
func PageData(db *sql.DB, posts []models.Post, style string, userId, page int) (models.MainPage, error) {
	var mainData models.MainPage
	pagination := NumberOfPages(len(posts))
	start, end := PostSlice(len(posts), page)
	posts = posts[start:end]
	var postData []models.PostData
	for _, post := range posts {
		data, err := GetPostData(db, post, userId)
		if err != nil {
			mainData.Metadata.Error = err.Error()
			return mainData, err
		}
		postData = append(postData, data)
	}
	var trendData []models.PostData
	if style == "main" || style == "filter" {
		trending, err := dbaser.TrendingPosts(db, 3)
		if err != nil {
			mainData.Metadata.Error = err.Error()
			return mainData, err
		}
		for _, post := range trending {
			data, err := GetPostData(db, post, userId)
			if err != nil {
				mainData.Metadata.Error = err.Error()
				return mainData, err
			}
			trendData = append(trendData, data)
		}
	}
	if style == "main" || style == "category" {
		categories, err := dbaser.Categories(db)
		if err != nil {
			mainData.Metadata.Error = err.Error()
			return mainData, err
		}
		mainData.Categories = categories
	}
	if userId > 0 {
		mainData.Metadata.LoggedIn = true
		userData, err := dbaser.UserById(db, userId)
		if err != nil {
			mainData.Metadata.Error = err.Error()
			return mainData, err
		}
		user := models.User{Avatar: userData.Avatar}
		mainData.User = user
	}
	pageData := models.Pagination{CurrentPage: page, TotalPages: pagination}
	mainData.Posts = postData
	mainData.Trending = trendData
	mainData.Pagination = pageData
	return mainData, nil
}

func UsersPageData(db *sql.DB, userId, userIdRequested, page int) (models.MainPage, error) {
	var mainData models.MainPage
	posts, err := dbaser.PostsByUser(db, userIdRequested)
	if err != nil {
		log.Print(err)
		mainData.Metadata.Error = err.Error()
		return mainData, err
	}
	mainData, err = PageData(db, posts, "main", userId, page)
	if err != nil {
		return mainData, err
	}
	return mainData, nil
}

func CategoryFilterPageData(db *sql.DB, categoryId, userId, page int) (models.MainPage, error) {
	var mainData models.MainPage
	posts, err := dbaser.PostsByCategory(db, categoryId, page)
	if err != nil {
		log.Print(err)
		mainData.Metadata.Error = err.Error()
		return mainData, err
	}
	mainData, err = PageData(db, posts, "category", userId, page)
	if err != nil {
		return mainData, err
	}
	return mainData, nil
}
