package config

import (
	"../structs"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	hostName := os.Getenv("HOST_DB_DEV")
	dbName := os.Getenv("NAME_DB_DEV")
	rootName := os.Getenv("USER_DB_DEV")
	dbPass := os.Getenv("PASS_DB_DEV")

	//db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/dashboard_go?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", rootName+":"+dbPass+"("+hostName+":3306)/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Person{})
	return db
}
