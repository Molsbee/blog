package service

import (
	"github.com/Molsbee/blog/model"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	id := session.Get("user_id")
	if id == nil {
		return false
	}

	return true
}

func SaveSessionData(user model.ServiceUser, c *gin.Context) error {
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Set("username", user.Username)
	session.Set("first_name", user.FirstName)
	session.Set("last_name", user.LastName)
	return session.Save()
}

func DeleteSessionData(c *gin.Context) error {
	session := sessions.Default(c)
	session.Clear()
	return session.Save()
}
