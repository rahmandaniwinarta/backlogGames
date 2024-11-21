package middlewares

import (
	"backlogGames/functions"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header Authorization
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Authorization token is required",
			})
			c.Abort()
			return
		}

		// Decode token JWT
		claims, err := functions.DecodeJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized: " + err.Error(),
			})
			c.Abort()
			return
		}

		fmt.Println("Decoded Role:", claims.Role)

		// Simpan username dan role ke context
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		// Lanjutkan ke handler berikutnya
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil role dari context
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": "Access restricted to admin only",
			})
			c.Abort()
			return
		}
		fmt.Println("Role in Context:", role)

		// Lanjutkan ke handler berikutnya
		c.Next()
	}
}
