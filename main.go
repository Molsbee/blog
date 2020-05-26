package main

import (
	"database/sql"
	"fmt"
	"github.com/Molsbee/blog/controller"
	"github.com/Molsbee/blog/handler"
	"github.com/Molsbee/blog/service"
	"github.com/gin-gonic/contrib/sessions"
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
	databaseURL := getEnvOrDefault("DATABASE_URL", fmt.Sprintf("postgres://%s:%s@%s:5432/blog?sslmode=disable", dbUsername, dbPassword, dbHostname))

	db, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		log.Panicf("failed to open connection to database - %s", err)
	}
	defer db.Close()
	runDatabaseMigration(db.DB())

	authService := service.NewAuthService(db)
	authController := controller.NewAuthController(authService)
	articleService := service.NewArticleService(db)
	articleController := controller.NewArticleController(articleService)

	router := gin.Default()
	// Setup Cookie Session
	router.Use(sessions.Sessions("user_session", sessions.NewCookieStore([]byte("secret"))))
	router.POST("/login", authController.Login)
	router.GET("/session", authController.Session)
	router.GET("/logout", authController.Logout)

	// Serve UI Static Content
	router.StaticFS("/css", http.Dir("./frontend/dist/css"))
	router.StaticFS("/img", http.Dir("./frontend/dist/img"))
	router.StaticFS("/js", http.Dir("./frontend/dist/js"))
	router.StaticFile("/favicon.ico", "./frontend/dist/favicon.ico")
	router.StaticFile("/", "./frontend/dist/index.html")
	router.GET("/admin/*subpage", handler.SessionRequiredHandler, func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})
	router.GET("/blog/*subpage", func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})

	// REST API with CORS Handler and Basic Authentication
	basicAuthHandler := handler.GetBasicAuthHandler(authService)
	api := router.Group("/api", handler.CORS)
	articles := api.Group("/articles")
	{
		articles.GET("", articleController.ListArticles)
		articles.GET("/:articleID", articleController.GetArticle)
		articles.POST("", basicAuthHandler, articleController.Create)
		articles.PUT("/:articleID", basicAuthHandler, articleController.UpdateArticle)
	}

	// serve web pages
	port := getEnvOrDefault("PORT", "8080")
	router.Run(fmt.Sprintf(":%s", port))
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
