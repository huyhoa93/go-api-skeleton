package news

type Post struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

type PostData struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []PostData `json:"data"`
}

type ResponseOne struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    PostData `json:"data"`
}
