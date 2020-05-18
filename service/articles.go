package service

import (
	"github.com/Molsbee/blog/model"
	"github.com/Molsbee/blog/repository"
	"github.com/jinzhu/gorm"
)

type ArticleService struct {
	articleRepository *repository.ArticleRepository
}

func NewArticleService(db *gorm.DB) *ArticleService {
	return &ArticleService{
		articleRepository: repository.NewArticleRepository(db),
	}
}

func (as *ArticleService) Create(request model.ArticleRequest) {

}

func (as *ArticleService) List() []model.ArticleResponse {
	return nil
}

func (as *ArticleService) Get(articleID int) model.ArticleResponse {
	return model.ArticleResponse{}
}

func (as *ArticleService) Update(articleID int, request model.ArticleRequest) {

}
