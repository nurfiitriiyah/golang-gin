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
		var (
			UserLogin []structs.TbUserLogins
		)
		idb.DB.Where(map[string]interface{}{"user_uname": user.Username, "user_password": user.Password}).Find(&UserLogin)
		if len(UserLogin) <= 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "wrong username or password",
			})
		} else {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": user.Username,
				"password": user.Password,
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
		//result = gin.H{
		//	"result": UserLogin,
		//}
		//

		//if user.Username != "myname" || user.Password != "myname123" {
		//	c.JSON(http.StatusUnauthorized, gin.H{
		//		"status":  http.StatusUnauthorized,
		//		"message": "wrong username or password",
		//	})
		//	c.Abort()
		//
		//} else {

		//	if err != nil {
		//		c.JSON(http.StatusInternalServerError, gin.H{
		//			"message": err.Error(),
		//		})
		//		c.Abort()
		//	}else {
		//		c.JSON(http.StatusOK, gin.H{
		//			"token": tokenString,
		//		})
		//	}
		//}

	}

}
