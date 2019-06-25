package controllers

import (
	"../structs"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
)

func (idb *InDB) CheckFirebase(c *gin.Context) {
	/*
		Declare variable, that will be used in this function
			var user contain struct of credential used for login
			var prepUserLogin contain struct of field in table tb_user_logins
			var UserLogin contain struct of field in table tb_user_logins and declared as array, for using len
	*/
	var user structs.UserFirebase

	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
		c.Abort()
	} else {
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

		type Post struct {
			Author       string `json:"author,omitempty"`
			Title        string `json:"title,omitempty"`
			Notification string `json:"notification"`
		}
		postsRef := ref.Child(user.UserID)
		if err := postsRef.Set(ctx, &Post{
			Notification: user.Notification,
		}); err != nil {
			log.Fatalln("Error setting value:", err)
		}

	}
}
