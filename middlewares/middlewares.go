package middlewares

import (
	"net/http"

	"FP-Sanbercode-Go-47-HaritsArkaanPutranto/utils/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		// Extract the user's role from the token
		role, err := token.ExtractUserRole(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		// Set the user's role in the context
		c.Set("user_role", role)

		c.Next()
	}
}
