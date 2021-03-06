package main

import (
	"./config"
	"./controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"time"

	//"google.golang.org/api/option"
	"log"
)

func main() {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:8080"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router.GET("/ots", inDB.GetOTS)
	router.POST("/checkFirebase", inDB.CheckFirebase)

	router.POST("/login", inDB.CheckLogin)
	router.POST("/detail/ots", inDB.GetDetailOTS)

	router.GET("/persons/:id", inDB.GetPersons)

	//router.GET("/person/:id", auth, inDB.GetPerson)
	//router.GET("/checkAuth", auth)
	//router.POST("/person", inDB.CreatePerson)
	//router.PUT("/person", auth, inDB.UpdatePerson)
	//router.DELETE("/person/:id", auth, inDB.DeletePerson)

	router.Run(":10005")
}
