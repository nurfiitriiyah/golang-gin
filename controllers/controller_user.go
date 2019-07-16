package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
)

var validate *validator.Validate

func (idb *InDB) CreateUser(c *gin.Context) {
	token, err := parseBearerToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, err.Error())
		c.Abort()
	} else {
		decoded := token.Claims
		fmt.Println("--------------------------")
		fmt.Println(decoded)

		//var createParams structs.CreateUser
		//var logins []structs.TbUserLogins

		//errs := c.BindJSON(&createParams)
		//if errs != nil {
		//	c.JSON(http.StatusUnauthorized, errs.Error())
		//	c.Abort()
		//} else {
		//	config := &validator.Config{TagName: "validate"}
		//	validate = validator.New(config)
		//	createUser := &structs.ValidateUser{
		//		UserName:     createParams.UserName,
		//		UserPassword: createParams.UserPassword,
		//		UserPlant:    createParams.UserPlant,
		//		UserRole:     createParams.UserRole,
		//		UserUname:    createParams.UserUname,
		//	}
		//	err := validate.Struct(createUser)
		//	if err != nil {
		//		c.JSON(http.StatusBadRequest, gin.H{
		//			"status":  http.StatusBadRequest,
		//			"message": "Paramater Incomplate",
		//		})
		//		c.Abort()
		//	} else {
		//		err := idb.DB.Where("user_uname = ?", createParams.UserUname).Find(&logins).Error
		//		if err != nil {
		//			fmt.Println(err)
		//		} else {
		//			if len(logins) < 1 {
		//				createIds := createIds("US")
		//				encrpt := encrypt(createParams.UserPassword)
		//				var Insert = structs.TbUserLogins{
		//					User_id:       createIds,
		//					User_name:     createParams.UserName,
		//					User_uname:    createParams.UserUname,
		//					User_password: encrpt,
		//					Role_id:       createParams.UserRole,
		//					User_plan:     createParams.UserPlant,
		//					User_status:   1,
		//				}
		//				idb.DB.Create(&Insert)
		//				c.JSON(http.StatusOK, gin.H{
		//					"status": "ok",
		//				})
		//			} else {
		//				c.JSON(http.StatusBadRequest, gin.H{
		//					"status":  http.StatusBadRequest,
		//					"message": "Username Already Exist",
		//				})
		//				c.Abort()
		//			}
		//		}
		//
		//	}
		//}
	}

}
