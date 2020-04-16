package comments

type Comments struct {
	Id      int    `json:"id,omitempty"`
	NewsId  int    `json:"news_id,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type CommentData struct {
	NewsId  int    `json:"news_id"  validate:"nonzero"`
	Comment string `json:"comment"  validate:"nonzero"`
}
