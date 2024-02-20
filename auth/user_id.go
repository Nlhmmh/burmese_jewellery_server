package auth

import "github.com/gin-gonic/gin"

const (
	userIDKey = "userID"
)

func SetUserID(c *gin.Context, userID string) {
	c.Set(userIDKey, userID)
}

func GetUserID(c *gin.Context) string {
	userID, exists := c.Get(userIDKey)
	if !exists {
		return ""
	}
	return userID.(string)
}
