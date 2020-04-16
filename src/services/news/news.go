package news

import (
	"net/http"

	connection "../../connection"
	comments "../../models/comments"
	news "../../models/news"
)

var newsTable string = "news"
var commentsTable string = "comments"

type response news.ResponsePosts

type responseOne news.ResponsePost

func GetNews(page, perpage int) response {
	db := connection.DBConn()
	defer db.Close()
	var total int
	result := db.Table(newsTable).Count(&total)
	if page != 0 && perpage != 0 {
		var offset int = (page - 1) * perpage
		result = result.Limit(perpage).Offset(offset)
	}
	var posts []news.PostData
	result = result.Order("id desc").Preload("Comments").Find(&posts)
	if result.Error != nil {
		res := response{
			Status:  http.StatusInternalServerError,
			Message: "Server Internal Error",
		}
		return res
	}
	res := response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    posts,
		Total:   total,
	}
	return res
}

func AddNews(title, content string) response {
	db := connection.DBConn()
	defer db.Close()
	var post = news.Post{
		Title:   title,
		Content: content,
	}
	result := db.Table(newsTable).Create(&post)
	if result.Error != nil {
		res := response{
			Status:  http.StatusInternalServerError,
			Message: "Server Internal Error",
		}
		return res
	}
	res := response{
		Status:  http.StatusCreated,
		Message: "Success",
	}
	return res
}

func GetNewsById(id int) responseOne {
	db := connection.DBConn()
	defer db.Close()
	var post news.PostData
	result := db.Table(newsTable).Preload("Comments").FirstOrInit(&post, id)
	if result.Error != nil {
		res := responseOne{
			Status:  http.StatusInternalServerError,
			Message: "Server Internal Error",
		}
		return res
	}
	if post.Id == 0 {
		res := responseOne{
			Status:  http.StatusNotFound,
			Message: "Not Found",
			Data:    post,
		}
		return res
	}
	res := responseOne{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    post,
	}
	return res
}

func UpdateNews(id int, title, content string) response {
	db := connection.DBConn()
	defer db.Close()
	var post news.PostData
	var data = news.Post{
		Title:   title,
		Content: content,
	}
	result := db.Table(newsTable).Find(&post, id).Updates(data)
	if result.Error != nil {
		res := response{
			Status:  http.StatusInternalServerError,
			Message: "Server Internal Error",
		}
		return res
	}
	res := response{
		Status:  http.StatusOK,
		Message: "Success",
	}
	return res
}

func DeleteNews(id int) response {
	db := connection.DBConn()
	defer db.Close()
	var post news.PostData
	result := db.Table(newsTable).Find(&post, id).Delete(&post)
	if result.Error != nil {
		res := response{
			Status:  http.StatusInternalServerError,
			Message: "Server Internal Error",
		}
		return res
	}
	res := response{
		Status:  http.StatusOK,
		Message: "Success",
	}
	return res
}

func AddComment(newsId int, comment string) response {
	db := connection.DBConn()
	defer db.Close()
	var commentData = comments.CommentData{
		NewsId:  newsId,
		Comment: comment,
	}
	result := db.Table(commentsTable).Create(&commentData)
	if result.Error != nil {
		res := response{
			Status:  http.StatusInternalServerError,
			Message: "Server Internal Error",
		}
		return res
	}
	res := response{
		Status:  http.StatusCreated,
		Message: "Success",
	}
	return res
}
