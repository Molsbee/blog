package repository

import (
	"github.com/Molsbee/blog/model"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type ArticleRepository interface {
	Save(article model.Article) model.ApplicationError
	FindAll() ([]model.Article, model.ApplicationError)
	FindByID(id int) (model.Article, model.ApplicationError)
}

type articleRepository struct {
	DB *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{
		DB: db,
	}
}

func (ar *articleRepository) Save(article model.Article) (err model.ApplicationError) {
	if dbErr := ar.DB.Create(&article).Error; dbErr != nil {
		log.Printf("unable to store article - %s\n", dbErr)
		err = model.ErrorBuilder().
			StatusCode(http.StatusInternalServerError).
			Message("unable to store article").
			Build()
	}

	return
}

func (ar *articleRepository) FindAll() (articles []model.Article, err model.ApplicationError) {
	dbErr := ar.DB.Find(&articles).Error
	if gorm.IsRecordNotFoundError(dbErr) {
		err = model.ErrorBuilder().StatusCode(http.StatusNotFound).Build()
	} else if dbErr != nil {
		log.Printf("unable to find all articles - %s\n", dbErr)
		err = model.ErrorBuilder().
			StatusCode(http.StatusInternalServerError).
			Message("unable to lookup all articles").
			Build()
	}

	return
}

func (ar *articleRepository) FindByID(id int) (article model.Article, err model.ApplicationError) {
	dbErr := ar.DB.Where("id = ?", id).Find(&article).Error
	if gorm.IsRecordNotFoundError(dbErr) {
		err = model.ErrorBuilder().StatusCode(http.StatusNotFound).Build()
	} else if dbErr != nil {
		log.Printf("failed to retrieve article - %s\n", dbErr)
		err = model.ErrorBuilder().
			StatusCode(http.StatusInternalServerError).
			Message("failed to retrieve article").
			Build()
	}

	return
}
