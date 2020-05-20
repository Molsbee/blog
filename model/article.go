package model

import "time"

type Article struct {
	ID        int       `gorm:"column:id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Published bool      `gorm:"column:published"`
	Title     string    `gorm:"column:title"`
	Content   string    `gorm:"column:content"`
	Author    string    `gorm:"column:author"`
}

func (Article) TableName() string {
	return "articles"
}

func (a Article) Description() string {
	if len(a.Content) < 100 {
		return a.Content
	}

	return a.Content[0:100]
}

type ArticleResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type ArticleRequest struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	Published bool   `json:"published"`
}
