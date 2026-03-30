package routes

import (
	"net/http"
	"rest_api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func getSingleUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "Couldn't parse user id",
		})

		return
	}

	userInfo, err := models.GetSingleUser(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Failed to get single user",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"user": userInfo,
	})
}

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

func loginUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "Couldn't parse requested data",
		})

		return
	}

	userInfo, err := user.Login(user.Email, user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Couldn't login user",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "logged in successfully",
		"user":    userInfo,
	})
}
