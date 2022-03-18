package auth

import (
	"net/http"
	"os"
	"time"

	common "go_api/src/models/common"
	auth "go_api/src/services/auth"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Create the JWT key used to create the signature
var jwtKey = []byte(os.Getenv("SECRET_KEY"))

// Create a struct to read the username and password from the request body
type Credentials common.Credentials

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims common.Claims

func Login(c *gin.Context) {
	var data Credentials
	if err := c.ShouldBindJSON(&data); err == nil {
		userId := auth.Login(data.Username, data.Password)
		if userId == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			return
		}

		// Declare the expiration time of the token
		// here, we have kept it as 14 days if exist param is_keep_loging
		var expiry_time int64
		expiry_time = time.Now().Unix() + 1*24*60*60
		if data.IsKeepLoging {
			expiry_time = time.Now().Unix() + 14*24*60*60
		}
		// Create the JWT claims, which includes the username and expiry time
		claims := &Claims{
			Id:       userId,
			Username: data.Username,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expiry_time,
			},
		}

		// Declare the token with the algorithm used for signing, and the claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Create the JWT string
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Server Internal Error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":      http.StatusOK,
			"token":       tokenString,
			"expiry_time": expiry_time,
		})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Bad Request",
		})
	}
}
