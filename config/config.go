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

	dbHost := os.Getenv("HOST_DB_DEV")
	dbName := os.Getenv("NAME_DB_DEV")
	dbRoot := os.Getenv("USER_DB_DEV")
	dbPass := os.Getenv("PASS_DB_DEV")

	db, err := gorm.Open("mysql", dbRoot+":"+dbPass+"("+dbHost+":3306)/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.DB().SetMaxIdleConns(100)

	db.AutoMigrate(structs.TbUserLogins{})
	db.AutoMigrate(structs.TbOutstanding{})
	db.AutoMigrate(structs.TbDelivery{})
	db.AutoMigrate(structs.TbLog{})
	db.AutoMigrate(structs.TbRetail{})
	return db
}
