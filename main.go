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
	log.Println("starting application by parsing environment variables")
	databaseURL := getEnvOrDefault("DATABASE_URL", "")
	if len(databaseURL) != 0 {
		// This is for used with Heroku which requires SSL for accessing Postgres
		databaseURL = fmt.Sprintf("%s?sslmode=require", databaseURL)
	} else {
		databaseURL = getEnvOrDefault("BLOG_DB_URL", "postgres://blogger:password@localhost:5432/blog?sslmode=disable")
	}

	log.Printf("setting up database connection %s\n", databaseURL)
	db, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		log.Panicf("failed to open connection to database - %s", err)
	}
	defer db.Close()
	runDatabaseMigration(db.DB())

	log.Println("setting up services and controllers")
	authService := service.NewAuthService(db)
	authController := controller.NewAuthController(authService)
	articleService := service.NewArticleService(db)
	articleController := controller.NewArticleController(articleService)

	log.Println("setting up router with paths")
	gin.SetMode(gin.ReleaseMode)
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
	port := fmt.Sprintf(":%s", getEnvOrDefault("PORT", "8080"))
	log.Printf("starting to listen %s\n", port)
	router.Run(port)
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

	log.Println("setting up new database migration with specified files")
	m, err := migrate.NewWithDatabaseInstance("file://database-migrations", "postgres", driver)
	if err != nil {
		log.Fatal("failed to setup the new database migrations", err)
	}

	log.Println("attempting to run database migrations")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run database migration scripts", err)
	}
}
