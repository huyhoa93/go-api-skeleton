package connection

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBConn() (db *gorm.DB) {
	dbDriver := "mysql"
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	db, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	return db
}
