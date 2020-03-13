package utility

import (
	"os"

	auth "../services/auth"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Create the JWT key used to create the signature
var jwtKey = []byte(os.Getenv("SECRET_KEY"))

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func ValidateToken(c *gin.Context) bool {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		return false
	} else {
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			return false
		}
		if !tkn.Valid {
			return false
		}
		userId := claims.Id
		check := auth.CheckUser(userId)
		if !check {
			return false
		}
		return true
	}
}
