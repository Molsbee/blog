package controller

import (
	"github.com/Molsbee/blog/service"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type authController struct {
	auth service.AuthService
}

func NewAuthController(authService service.AuthService) *authController {
	return &authController{
		auth: authService,
	}
}

func (auth *authController) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameters cannot be empty"})
		return
	}

	user, err := auth.auth.Authenticate(username, password)
	if err != nil {
		c.JSON(err.StatusCode(), err.Details())
	}

	if err := service.SaveSessionData(*user, c); err != nil {
		log.Printf("saving session failed - %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save session"})
		return
	}

	c.JSON(200, gin.H{
		"username": user.Username,
	})
}

func (auth *authController) Session(c *gin.Context) {
	if service.HasSession(c) {
		session := sessions.Default(c)
		c.JSON(200, gin.H{
			"username": session.Get("username"),
		})
		return
	}

	c.Status(http.StatusNotFound)
}

func (auth *authController) Logout(c *gin.Context) {
	hasSession := service.HasSession(c)
	if hasSession {
		if err := service.DeleteSessionData(c); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save session"})
			return
		}
	}

	c.Status(http.StatusOK)
}
