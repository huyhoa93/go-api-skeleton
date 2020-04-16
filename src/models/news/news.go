package news

import (
	comments "../comments"
)

type Post struct {
	Title   string `form:"title" json:"title" binding:"required" validate:"nonzero,min=1,regexp=^[a-zA-Z0-9]*$"`
	Content string `form:"content" json:"content" binding:"required" validate:"nonzero"`
}

type PostData struct {
	Id       int                 `json:"id,omitempty"`
	Title    string              `json:"title,omitempty"`
	Content  string              `json:"content,omitempty"`
	Comments []comments.Comments `json:"comments,omitempty" gorm:"foreignkey:NewsId;association_foreignkey:Id"`
}

type ResponsePosts struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []PostData `json:"data,omitempty"`
	Total   int        `json:"total,omitempty"`
}

type ResponsePost struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    PostData `json:"data,omitempty"`
}
