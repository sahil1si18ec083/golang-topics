package middlewares

import (
	"fmt"
	"net/http"
	"rest-api/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header missing",
			})
			c.Abort()
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Authorization format",
			})
			c.Abort()
			return

		}
		authToken := parts[1]
		fmt.Println(authHeader, "jjj")
		fmt.Println("Dot Count =", strings.Count(authToken, "."))
		claims, err := utils.VerifyToken(authToken)

		fmt.Println(err)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "UnAuthorized",
			})
			c.Abort()
			return
		}
		c.Set("userId", claims.ID)
		c.Set("email", claims.Email)
		c.Next()

	}
}
