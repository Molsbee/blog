package handler

import (
	"github.com/Molsbee/blog/repository"
	"github.com/Molsbee/blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBasicAuthHandler(repository repository.ServiceUserRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		user := repository.FindByUsernameAndPassword(username, password)
		if user == nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}

func SessionRequiredHandler(c *gin.Context) {
	hasSession := service.HasSession(c)
	if !hasSession {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()
}
