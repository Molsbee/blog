package handler

import (
	"github.com/Molsbee/blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBasicAuthHandler(authService service.AuthService) func(c *gin.Context) {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		_, err := authService.Authenticate(username, password)
		if err != nil {
			c.AbortWithStatusJSON(err.StatusCode(), err.Details())
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
