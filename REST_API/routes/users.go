package routes

import (
	"net/http"
	"rest_api/models"
	"rest_api/utils"
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

	err = user.Login()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": err.Error(),
		})

		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Couldn't generate token",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "logged in successfully",
		"token":   token,
	})
}

func updateUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "Couldn't parse user id",
		})

		return
	}

	var user models.User
	err = c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "Couldn't parse user data",
		})

		return
	}

	_, err = models.GetSingleUser(userId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"ok":      false,
			"message": "Couldn't find user",
		})

		return
	}

	err = user.Update(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Couldn't update user info",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "User updated successfully",
	})
}

func deleteUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "Couldn't parse user id",
		})

		return
	}

	user, err := models.GetSingleUser(userId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"ok":      false,
			"message": "Couldn't find user",
		})

		return
	}

	err = user.Delete(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Couldn't update user info",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "User deleted successfully",
	})
}
