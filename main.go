package main

import (
	"github.com/Molsbee/blog/controller"
	"github.com/Molsbee/blog/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func main() {
	db, err := gorm.Open("mysql", "root:blog-development@blog_mysql-local-development_1:33060/blog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panicf("failed to open connection to database - %s", err)
	}
	defer db.Close()

	articleService := service.NewArticleService(db)
	articleController := controller.NewArticleController(articleService)

	router := gin.Default()
	articles := router.Group("/articles")
	{
		articles.POST("", articleController.Create)
		articles.GET("", articleController.ListArticles)
		articles.GET("/:articleID", articleController.GetArticle)
		articles.PUT("/:articleID", articleController.UpdateArticle)
	}

	// serve web pages
	router.Run(":8080")
}
