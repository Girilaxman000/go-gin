package middleware

import (
	"net/http"

	"github.com/Girilaxman000/go-gin/commonfunctions"
	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
)

type MyCustomClaims struct {
	Sub      string `json:"sub"`
	FullName string `json:"fullName"`
	jwt.StandardClaims
}

func AuthMiddleware(c *gin.Context) {

	tokenString, err := commonfunctions.ExtractAuthBearerToken(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort() // Stop further processing of the request
		return
	}

	token, err := commonfunctions.ParseClaims(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort() // Stop further processing of the request
		return
	}

	claims, ok := commonfunctions.RetrieveClaims(token)
	if ok && token.Valid {
		c.JSON(200, gin.H{
			"message": "Access granted",
			"sub":     claims.Sub,
		})
		c.Next()
	} else {
		c.String(401, "Unauthorized")
		c.Abort()
		return
	}

}
