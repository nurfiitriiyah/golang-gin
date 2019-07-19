package controllers

import (
	"../structs"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// to get one data with {id}
func (idb *InDB) CheckLogin(c *gin.Context) {
	/*
		Declare variable, that will be used in this function
			var user contain struct of credential used for login
			var prepUserLogin contain struct of field in table tb_user_logins
			var UserLogin contain struct of field in table tb_user_logins and declared as array, for using len
	*/
	var user structs.Credential
	var PrepUserLogin structs.TbUserLogins
	var UserLogin []structs.TbUserLogins
	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
		c.Abort()
	} else {
		if user.Username == "" || user.Password == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Parameter Incomplate",
			})
		} else {
			username := user.Username
			password := encrypt(user.Password)
			idb.DB.Where(structs.TbUserLogins{User_uname: username, User_password: password}).First(&UserLogin).Scan(&PrepUserLogin)
			if len(UserLogin) <= 0 {
				c.JSON(http.StatusUnauthorized, gin.H{
					"status":  http.StatusUnauthorized,
					"message": "Username or Password is not match",
				})
			} else {
				claimMap := jwt.MapClaims{
					"id":       PrepUserLogin.User_id,
					"name":     PrepUserLogin.User_name,
					"role":     PrepUserLogin.Role_id,
					"plan":     PrepUserLogin.User_plan,
					"username": PrepUserLogin.User_uname,
				}
				secret := os.Getenv("JWT_SECRET_KEY")

				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimMap)
				tokenString, err := token.SignedString([]byte(secret))
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"message": err.Error(),
					})
					c.Abort()
				} else {
					c.JSON(http.StatusOK, gin.H{
						"token": tokenString,
					})
				}
			}
		}

	}
}
