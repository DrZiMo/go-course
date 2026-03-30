package routes

import (
	"net/http"
	"rest_api/models"

	"github.com/gin-gonic/gin"
)

func createUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": err.Error(),
		})

		return
	}

	err = user.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Couldn't save user",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "User created successfully",
		"user": map[string]interface{}{
			"email":    user.Email,
			"password": user.Password,
		},
	})
}

func getUsers(c *gin.Context) {
	users, err := models.GetUsers()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "Couldn't fetch users",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":    true,
		"users": users,
	})
}
