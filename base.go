package main

import (
	"./config"
	"./controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	//"google.golang.org/api/option"
	"log"
)

const (
	serverKey = "YOUR-KEY"
	topic     = "/topics/someTopic"
)

func main() {

	router := gin.Default()
	router.Use(cors.Default())
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

	//router.GET("/person/:id", auth, inDB.GetPerson)
	//router.GET("/persons", auth, inDB.GetPersons)
	//router.GET("/checkAuth", auth)
	//router.POST("/person", inDB.CreatePerson)
	//router.PUT("/person", auth, inDB.UpdatePerson)
	//router.DELETE("/person/:id", auth, inDB.DeletePerson)

	router.Run(":10005")
}
