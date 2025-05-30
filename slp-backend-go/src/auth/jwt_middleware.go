package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Brak tokenu autoryzacyjnego"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := ValidateJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Nieprawidłowy token"})
			return
		}

		email, _ := claims["email"].(string)
		role, _ := claims["role"].(string)

		c.Set("email", email)
		c.Set("role", role)

		c.Next()
	}
}

func RequireMinRole(minRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !HasMinRole(c, minRole) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Brak uprawnień"})
			return
		}
		c.Next()
	}
}

var rolePriority = map[string]int{
	"INTERN": 0,
	"WORKER": 1,
	"ADMIN":  2,
}

func HasMinRole(c *gin.Context, required string) bool {
	roleVal, exists := c.Get("role")
	if !exists {
		return false
	}
	userRole := strings.ToUpper(roleVal.(string))
	return rolePriority[userRole] >= rolePriority[required]
}
