package main

import (
	"database/sql"
	"github.com/Molsbee/blog/controller"
	"github.com/Molsbee/blog/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func main() {
	db, err := gorm.Open("mysql", "root:blog-development@blog_mysql-local-development_1:3306/blog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panicf("failed to open connection to database - %s", err)
	}
	defer db.Close()
	runDatabaseMigration(db.DB())

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

func runDatabaseMigration(db *sql.DB) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://database-migrations", "mysql", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
