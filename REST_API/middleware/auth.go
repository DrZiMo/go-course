package middleware

import (
	"net/http"
	"rest_api/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"ok":      false,
			"message": "Not authorized",
		})

		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"ok":      false,
			"message": "Not authorized",
		})

		return
	}

	c.Set("userId", userId)

	c.Next()
}
