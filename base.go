package main

import (
	"./config"
	"./controllers"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strings"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {

	router := gin.Default()
	router.Use(cors.Default())

	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	router.POST("/login", inDB.CheckLogin)
	router.GET("/ots", inDB.GetOTS)
	router.POST("/detail/ots", inDB.GetDetailOTS)

	router.GET("/person/:id", auth, inDB.GetPerson)
	router.GET("/persons", auth, inDB.GetPersons)
	router.GET("/checkAuth", auth)
	router.POST("/person", inDB.CreatePerson)
	router.PUT("/person", auth, inDB.UpdatePerson)
	router.DELETE("/person/:id", auth, inDB.DeletePerson)

	router.Run(":10005")
}

func auth(c *gin.Context) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, " ")
		fmt.Println("---------------------------CHECK AUTH--------------------------------")
		token, err := parseBearerToken(bearerToken[0])
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
		} else {
			decoded := token.Claims
			c.JSON(http.StatusOK, decoded)
		}

	}
}

func parseBearerToken(bearerToken string) (*jwt.Token, error) {
	return jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("error")
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})

}
