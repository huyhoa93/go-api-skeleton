package common

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// Create a struct to read the username and password from the request body
type Credentials struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	IsKeepLoging bool   `json:"is_keep_loging"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type PaginationParams struct {
	Page    int `form:"page"`
	Perpage int `form:"perpage"`
}
