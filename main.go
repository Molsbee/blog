package main

import (
	"database/sql"
	"github.com/Molsbee/blog/controller"
	"github.com/Molsbee/blog/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
)

func main() {
	db, err := gorm.Open("postgres", "postgres://blog:blog-development@localhost:5432/blog?sslmode=disable")
	if err != nil {
		log.Panicf("failed to open connection to database - %s", err)
	}
	defer db.Close()
	runDatabaseMigration(db.DB())

	articleService := service.NewArticleService(db)
	articleController := controller.NewArticleController(articleService)

	router := gin.Default()
	// Serve Static Content
	router.StaticFS("/css", http.Dir("./frontend/dist/css"))
	router.StaticFS("/img", http.Dir("./frontend/dist/img"))
	router.StaticFS("/js", http.Dir("./frontend/dist/js"))
	router.StaticFile("/favicon.ico", "./frontend/dist/index.html")
	router.StaticFile("/", "./frontend/dist/index.html")

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
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://database-migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
