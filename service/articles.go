package service

import (
	"github.com/Molsbee/blog/model"
	"github.com/Molsbee/blog/repository"
	"github.com/jinzhu/gorm"
	"time"
)

type ArticleService struct {
	articleRepository repository.ArticleRepository
}

func NewArticleService(db *gorm.DB) *ArticleService {
	return &ArticleService{
		articleRepository: repository.NewArticleRepository(db),
	}
}

func (as *ArticleService) Create(request model.ArticleRequest) model.ApplicationError {
	return as.articleRepository.Save(model.Article{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Published: request.Published,
		Title:     request.Title,
		Content:   request.Content,
		Author:    request.Author,
	})
}

func (as *ArticleService) List() (articles []model.Article, err model.ApplicationError) {
	// Still need to order list so newer articles are first
	// and add additional support for query params to filter or limit content
	return as.articleRepository.FindAll()
}

func (as *ArticleService) Get(articleID int) (article model.Article, err model.ApplicationError) {
	return as.articleRepository.FindByID(articleID)
}

func (as *ArticleService) Update(articleID int, request model.ArticleRequest) model.ApplicationError {
	article, err := as.articleRepository.FindByID(articleID)
	if err != nil {
		return err
	}

	article.Author = request.Author
	article.Content = request.Content
	article.Title = request.Title
	article.Published = request.Published
	return as.articleRepository.Save(article)
}
