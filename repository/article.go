package repository

import (
	"fmt"
	"github.com/Molsbee/blog/model"
	"github.com/jinzhu/gorm"
	"log"
)

type ArticleRepository interface {
	Save(article model.Article) error
	FindAll() ([]model.Article, error)
	FindByID(id int) (model.Article, error)
}

type articleRepository struct {
	DB *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{
		DB: db,
	}
}

func (ar *articleRepository) Save(article model.Article) (err error) {
	if dbErr := ar.DB.Create(&article).Error; dbErr != nil {
		log.Printf("unable to store article - %s\n", dbErr)
		err = fmt.Errorf("unable to store article")
	}

	return
}

func (ar *articleRepository) FindAll() (articles []model.Article, err error) {
	if dbErr := ar.DB.Find(&articles).Error; dbErr != nil {
		log.Printf("unable to find all articles - %s\n", dbErr)
		err = fmt.Errorf("unable to look up all articles")
	}

	return
}

func (ar *articleRepository) FindByID(id int) (article model.Article, err error) {
	if dbErr := ar.DB.Where("id = ?", id).Find(&article).Error; dbErr != nil {
		log.Printf("failed to retrieve article - %s\n", dbErr)
		err = fmt.Errorf("failed to retrieve article")
	}

	return
}
