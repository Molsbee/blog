package controller

import (
	"github.com/Molsbee/blog/repository"
	"github.com/Molsbee/blog/service"
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

	if err := service.SaveSessionData(*user, c); err != nil {
		log.Printf("saving session failed - %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save session"})
		return
	}

	c.Status(200)
}

func (auth *authController) Logout(c *gin.Context) {
	hasSession := service.HasSession(c)
	if hasSession {
		if err := service.DeleteSessionData(c); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save session"})
			return
		}
	}

	c.Status(200)
}
