package utility

import (
	"os"

	common "go_api/src/models/common"
	auth "go_api/src/services/auth"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Create the JWT key used to create the signature
var jwtKey = []byte(os.Getenv("SECRET_KEY"))

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims common.Claims

func ValidateToken(c *gin.Context) bool {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		return false
	} else {
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !tkn.Valid {
			return false
		}
		check := auth.CheckUser(claims.Id)
		if !check {
			return false
		}
		return true
	}
}
