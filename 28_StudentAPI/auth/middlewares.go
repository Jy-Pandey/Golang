package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AUthMiddleware func return a function because Gin requires middleware to be func(c *gin.Context)
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 1. Get the "Authorization" header
		// authHeader := c.GetHeader("Authorization")

		// if authHeader == "" {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		// 	return
		// }

		// 2. Extract the token value
		// Token should start with "Bearer <token val>"
		// parts := strings.Split(authHeader, " ")
		// if len(parts) != 2 || parts[0] != "Bearer" {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		// 	return
		// }

		// tokenStr := parts[1];

		//* Using cookies
		tokenStr, err := c.Cookie("token")
		if err != nil || tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token in cookie"})
			return
		}

		// 3. Validate and decode token into Claims struct
		token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// 4. Extract the user info from the token Claims
		//* token.Claims is of type jwt.Claims (interface)
		//* But your real claims are of type *Claims (your struct)
		//* So you need to type assert it like this:
		claims, ok := token.Claims.(*Claims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Couldn't parse claims"})
			return
		}

		// 5. Set user info in Gin context for further use
		c.Set("userID", claims.UserId)
		c.Set("email", claims.Email)

		c.Next() // proceed to the actual route handler
	}
}
