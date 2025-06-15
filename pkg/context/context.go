package context

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const UserIDKey = "userID"

func SetUserID(c *gin.Context, userID uint) {
	c.Set(UserIDKey, userID)
}

func GetUserID(c *gin.Context) uint {
	if userID, exists := c.Get(UserIDKey); exists {
		if id, ok := userID.(uint); ok {
			return id
		}
	}
	return 0
}

func MustGetUserID(c *gin.Context) (uint, error) {
	if userID, exists := c.Get(UserIDKey); exists {
		if id, ok := userID.(uint); ok {
			return id, nil
		}
	}
	return 0, errors.New("userID not found in context")
}
