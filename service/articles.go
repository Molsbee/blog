package service

import (
	"github.com/Molsbee/blog/model"
	"github.com/Molsbee/blog/repository"
	"github.com/Molsbee/blog/repository/database_model"
	"github.com/jinzhu/gorm"
	"time"
)

type ArticleService struct {
	articleRepository *repository.ArticleRepository
}

func NewArticleService(db *gorm.DB) *ArticleService {
	return &ArticleService{
		articleRepository: repository.NewArticleRepository(db),
	}
}

func (as *ArticleService) Create(request model.ArticleRequest) error {
	return as.articleRepository.Save(database_model.Article{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Published: request.Published,
		Title:     request.Title,
		Content:   request.Content,
		Author:    request.Author,
	})
}

func (as *ArticleService) List() ([]map[string]interface{}, error) {
	articles, err := as.articleRepository.FindAll()
	if err != nil {
		return nil, err
	}

	articlesResponse := make([]map[string]interface{}, len(articles))
	for i, article := range articles {
		articlesResponse[i] = map[string]interface{}{
			"id":          article.ID,
			"title":       article.Title,
			"author":      article.Author,
			"description": article.Description(),
		}
	}

	return articlesResponse, nil
}

func (as *ArticleService) Get(articleID int) (model.ArticleResponse, error) {
	art, err := as.articleRepository.FindByID(articleID)
	if err != nil {
		return model.ArticleResponse{}, err
	}

	return model.ArticleResponse{
		ID:      art.ID,
		Title:   art.Title,
		Content: art.Content,
		Author:  art.Author,
	}, nil
}

func (as *ArticleService) Update(articleID int, request model.ArticleRequest) {

}
