package news

import (
	"net/http"

	connection "../connection"
)

type post struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []post `json:"data"`
}

type responseOne struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    post   `json:"data"`
}

func GetNews() response {
	db := connection.DBConn()
	rows, err := db.Query("SELECT id, title, content FROM news")
	var posts []post
	if err != nil {
		res := response{
			Status:  http.StatusInternalServerError,
			Message: "Server Internal Error",
		}
		return res
		// log.Fatal(err)
	}

	defer rows.Close()
	var p post
	for rows.Next() {
		var id int
		var title, content string
		err := rows.Scan(&id, &title, &content)
		if err != nil {
			res := response{
				Status:  http.StatusInternalServerError,
				Message: "Server Internal Error",
			}
			return res
			// log.Fatal(err)
		}
		p.Id = id
		p.Title = title
		p.Content = content
		posts = append(posts, p)
	}
	if err := rows.Err(); err != nil {
		res := response{
			Status:  http.StatusInternalServerError,
			Message: "Server Internal Error",
		}
		return res
		// log.Fatal(err)
	}
	res := response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    posts,
	}
	defer db.Close()
	return res
}

func AddNews(title, content string) response {
	db := connection.DBConn()
	_, err := db.Exec(`INSERT INTO news (title, content) VALUES (?, ?)`, title, content)
	if err != nil {
		res := response{
			Status:  http.StatusInternalServerError,
			Message: "Server Internal Error",
		}
		return res
		// log.Fatal(err)
	}
	res := response{
		Status:  http.StatusCreated,
		Message: "Success",
	}
	defer db.Close()
	return res
}

func GetNewsById(id int) responseOne {
	db := connection.DBConn()
	q := "SELECT title, content FROM news WHERE id = ?"
	var title, content string
	var post post
	if err := db.QueryRow(q, id).Scan(&title, &content); err != nil {
		res := responseOne{
			Status:  http.StatusInternalServerError,
			Message: "Server Internal Error",
		}
		return res
		// log.Fatal(err)
	}
	post.Id = id
	post.Title = title
	post.Content = content
	res := responseOne{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    post,
	}
	defer db.Close()
	return res
}

func UpdateNews(id int, title, content string) response {
	db := connection.DBConn()
	_, err := db.Exec(`UPDATE news SET title=?, content=? WHERE id=?`, title, content, id)
	if err != nil {
		res := response{
			Status:  http.StatusInternalServerError,
			Message: "Server Internal Error",
		}
		return res
		// log.Fatal(err)
	}
	res := response{
		Status:  http.StatusOK,
		Message: "Success",
	}
	defer db.Close()
	return res
}

func DeleteNews(id int) response {
	db := connection.DBConn()
	_, err := db.Exec(`DELETE FROM news WHERE id = ?`, id)
	if err != nil {
		res := response{
			Status:  http.StatusInternalServerError,
			Message: "Server Internal Error",
		}
		return res
		// log.Fatal(err)
	}
	res := response{
		Status:  http.StatusOK,
		Message: "Success",
	}
	defer db.Close()
	return res
}
