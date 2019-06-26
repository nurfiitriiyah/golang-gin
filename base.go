package main

import (
	"./config"
	"./controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	hostFront := os.Getenv("FRONTEND")
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{hostFront},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router.POST("/checkFirebase", inDB.CheckFirebase)

	router.POST("/login", inDB.CheckLogin)

	router.GET("/ots", inDB.GetOTS)
	router.POST("/detail/ots", inDB.GetDetailOTS)

	router.GET("/ios", inDB.GetIOS)
	//router.POST("/detail/ots", inDB.GetDetailOTS)

	router.Run(":10005")
}
