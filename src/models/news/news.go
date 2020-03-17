package news

type Post struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

type PostData struct {
	Id      int    `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type ResponsePosts struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []PostData `json:"data,omitempty"`
}

type ResponsePost struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    PostData `json:"data,omitempty"`
}
