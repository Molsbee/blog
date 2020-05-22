package controller

import (
	"github.com/Molsbee/blog/repository"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type authController struct {
	repo repository.ServiceUserRepository
}

func NewAuthController(userRepository repository.ServiceUserRepository) *authController {
	return &authController{
		repo: userRepository,
	}
}

func (auth *authController) Login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameters cannot be empty"})
		return
	}

	user := auth.repo.FindByUsernameAndPassword(username, password)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication failed"})
		return
	}

	session.Set("user", user.Username)
	if err := session.Save(); err != nil {
		log.Printf("saving session failed - %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save session"})
		return
	}

	c.Status(200)
}

func (auth *authController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session token"})
		return
	}

	session.Delete("user")
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save session"})
		return
	}

	c.Status(200)
}
