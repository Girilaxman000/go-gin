package commonfunctions

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Girilaxman000/go-gin/api/dtos"
	"github.com/Girilaxman000/go-gin/api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type MyCustomClaims struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func HashPassword(databaseUser models.User, user dtos.UserLoginDetail) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(databaseUser.Password), []byte(user.Password))
	return err
}

func GenerateToken(databaseUser models.User) (tokenString string, err error) {
	//this is payload in jwt token
	claims := MyCustomClaims{
		Sub:   strconv.FormatUint(uint64(databaseUser.ID), 10),
		Email: databaseUser.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString([]byte(os.Getenv("SECRET")))
	return tokenString, err
}

func ExtractAuthBearerToken(c *gin.Context) (token string, err error) {
	header := c.GetHeader("Authorization")
	if header == "" {
		return "", errors.New("Token not found")
	}

	if !strings.Contains(header, "Bearer") {
		return "", errors.New("Bearer token not found")
	}
	tokenString := strings.TrimSpace(strings.Replace(header, "Bearer", "", 1))
	return tokenString, nil
}

func ParseClaims(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	return token, err
	// here written type must be token address of
	// return &token, err look at it
}

func RetrieveClaims(token *jwt.Token) (*MyCustomClaims, bool) {
	// Verfify token
	claims, ok := token.Claims.(*MyCustomClaims)

	return claims, ok

}
