package auth

import (
	"crypto/md5"
	"encoding/hex"

	connection "go_api/src/connection"
	users "go_api/src/models/users"
)

var usersTable string = "users"

func Login(username, password string) int {
	db := connection.DBConn()
	defer db.Close()
	hasher := md5.New()
	hasher.Write([]byte(password))
	pass := hex.EncodeToString(hasher.Sum(nil))
	var user users.User
	result := db.Table(usersTable).Select("id").FirstOrInit(&user, "username = ? AND password = ?", username, pass)
	if result.Error != nil {
		return 0
	}
	return user.Id
}

func CheckUser(id int) bool {
	db := connection.DBConn()
	defer db.Close()
	var user users.User
	result := db.Table(usersTable).Select("username").FirstOrInit(&user, id)
	if result.Error != nil {
		return false
	}
	if user.Username == "" {
		return false
	}
	return true
}
