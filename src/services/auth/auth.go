package auth

import (
	"crypto/md5"
	"encoding/hex"

	connection "../../connection"
)

func Login(username, password string) int {
	db := connection.DBConn()
	hasher := md5.New()
	hasher.Write([]byte(password))
	pass := hex.EncodeToString(hasher.Sum(nil))
	var id int
	if err := db.QueryRow(`SELECT id FROM users WHERE username=? AND password=?`, username, pass).Scan(&id); err != nil {
		return 0
	}
	return id
}

func CheckUser(id int) bool {
	db := connection.DBConn()
	var username string
	if err := db.QueryRow(`SELECT username FROM users WHERE id=?`, id).Scan(&username); err != nil {
		return false
	}
	if username == "" {
		return false
	}
	return true
}
