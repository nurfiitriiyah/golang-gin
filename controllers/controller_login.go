package controllers

import (
	"../structs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

// to get one data with {id}
func (idb *InDB) CheckLogin(c *gin.Context) {
	var user structs.Credential
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
		c.Abort()
	} else {
		var PrepUserLogin structs.TbUserLogins
		var UserLogin []structs.TbUserLogins

		idb.DB.Where(structs.TbUserLogins{User_uname: user.Username, User_password: user.Password}).First(&UserLogin).Scan(&PrepUserLogin)

		if len(UserLogin) <= 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "wrong username or password",
			})
		} else {

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"id":       PrepUserLogin.User_id,
				"name":     PrepUserLogin.User_name,
				"plan":     PrepUserLogin.User_plan,
				"username": PrepUserLogin.User_uname,
			})
			tokenString, err := token.SignedString([]byte("secret"))
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
