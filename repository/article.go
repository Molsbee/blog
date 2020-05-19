package repository

import (
	"fmt"
	"github.com/Molsbee/blog/repository/database_model"
	"github.com/jinzhu/gorm"
	"log"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{
		DB: db,
	}
}

func (ar *ArticleRepository) Save(article database_model.Article) (err error) {
	if dbErr := ar.DB.Create(&article).Error; dbErr != nil {
		log.Printf("unable to store article - %s\n", dbErr)
		err = fmt.Errorf("unable to store article")
	}

	return
}

func (ar *ArticleRepository) FindAll() (articles []database_model.Article, err error) {
	if dbErr := ar.DB.Find(&articles).Error; dbErr != nil {
		log.Printf("unable to find all articles - %s\n", dbErr)
		err = fmt.Errorf("unable to look up all articles")
	}

	return
}

func (ar *ArticleRepository) FindByID(id int) (article database_model.Article, err error) {
	if dbErr := ar.DB.Where("id = ?", id).Find(&article).Error; dbErr != nil {
		log.Printf("failed to retrieve article - %s\n", dbErr)
		err = fmt.Errorf("failed to retrieve article")
	}

	return
}
