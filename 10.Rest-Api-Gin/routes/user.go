package routes

import (
	"fmt"
	"net/http"
	"rest-api/models"
	"rest-api/utils"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var req models.User
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	hashedPassword, err := utils.HashPssword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.Password = hashedPassword

	err = req.Save()
	fmt.Println(req)
	fmt.Println("family man")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create user",
		})
		return
	}
	token, err := utils.GenerateJWT(req.Email, req.ID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create user",
		})
		return
	}
	fmt.Println(token)

	c.JSON(http.StatusOK, gin.H{
		"message": "Signup successfully",
		"token":   token,
	})

}
func Login(c *gin.Context) {
	var req models.User
	var user models.User
	err := c.BindJSON(&req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}
	err = user.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	err = utils.ComparePassword(user.Password, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	token, err := utils.GenerateJWT(user.Email, user.ID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create user",
		})
		return
	}
	fmt.Println(token)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successfully",
		"token":   token,
	})

}
