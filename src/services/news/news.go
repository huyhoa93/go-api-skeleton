package news

import (
	"net/http"

	connection "../../connection"
	news "../../models/news"
)

var newsTable string = "news"

type response news.ResponsePosts

type responseOne news.ResponsePost

func GetNews() response {
	db := connection.DBConn()
	defer db.Close()
	var posts []news.PostData
	result := db.Table(newsTable).Find(&posts)
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
	result := db.Table(newsTable).FirstOrInit(&post, id)
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
