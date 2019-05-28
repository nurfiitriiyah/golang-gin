package controllers

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"log"
	"os"
)

func (idb *InDB) CheckFirebase(c *gin.Context) {
	fmt.Println("-------------------------------------FIREBASE----------------------------------------------")

	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile("refreshToken.json")

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	//As an admin, the app has access to read and write all data, regradless of Security Rules
	ref := client.NewRef("User")
	var data map[string]interface{}
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}

	type User struct {
		DateOfBirth string `json:"date_of_birth,omitempty"`
		FullName    string `json:"full_name,omitempty"`
		Nickname    string `json:"nickname,omitempty"`
	}

	a := ref.Set(ctx, map[string]*User{
		"alanisawesome": {
			DateOfBirth: "June 23, 1912",
			FullName:    "Alan Turing",
		},
		"gracehop": {
			DateOfBirth: "December 9, 1906",
			FullName:    "Grace Hopper",
		},
	})
	if a != nil {
		log.Fatalln("Error setting value:", err)
	}
}
