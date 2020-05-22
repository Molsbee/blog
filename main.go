package main

import (
	"database/sql"
	"fmt"
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
	"os"
)

func main() {
	dbHostname := getEnvOrDefault("BLOG_DATABASE_HOSTNAME", "localhost")
	dbUsername := getEnvOrDefault("BLOG_DATABASE_USERNAME", "blog")
	dbPassword := getEnvOrDefault("BLOG_DATABASE_PASSWORD", "blog-development")

	db, err := gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:5432/blog?sslmode=disable", dbUsername, dbPassword, dbHostname))
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
	router.StaticFile("/favicon.ico", "./frontend/dist/favicon.ico")
	router.StaticFile("/", "./frontend/dist/index.html")

	// Setup CORS Handler and Authorization Handler
	api := router.Group("/api", corsHandler)
	articles := api.Group("/articles")
	{
		articles.POST("", articleController.Create)
		articles.GET("", articleController.ListArticles)
		articles.GET("/:articleID", articleController.GetArticle)
		articles.PUT("/:articleID", articleController.UpdateArticle)
	}

	// serve web pages
	router.Run(":8080")
}

func getEnvOrDefault(environmentVariable string, defaultValue string) string {
	variable := os.Getenv(environmentVariable)
	if len(variable) == 0 {
		return defaultValue
	}

	return variable
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

func corsHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
