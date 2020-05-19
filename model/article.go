package model

type ArticleResponse struct {
	ID      int
	Title   string
	Content string
	Author  string
}

type ArticleRequest struct {
	Title     string
	Content   string
	Author    string
	Published bool
}
