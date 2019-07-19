package controllers

import (
	"../structs"
	"fmt"
	"github.com/dgrijalva/jwt-go"
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
		decoded := token.Claims.(jwt.MapClaims)
		id := fmt.Sprintf("%v", decoded["id"])
		var createParams structs.CreateUser
		var logins []structs.TbUserLogins

		errs := c.BindJSON(&createParams)
		if errs != nil {
			c.JSON(http.StatusBadRequest, errs.Error())
			c.Abort()
		} else {
			config := &validator.Config{TagName: "validate"}
			validate = validator.New(config)
			createUser := &structs.ValidateUser{
				UserName:     createParams.UserName,
				UserPassword: createParams.UserPassword,
				UserPlant:    createParams.UserPlant,
				UserRole:     createParams.UserRole,
				UserUname:    createParams.UserUname,
			}
			err := validate.Struct(createUser)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  http.StatusBadRequest,
					"message": "Paramater Incomplate",
				})
				c.Abort()
			} else {
				err := idb.DB.Where("user_uname = ?", createParams.UserUname).Find(&logins).Error
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  http.StatusBadRequest,
						"message": "Paramater Incomplate",
					})
				} else {
					if len(logins) < 1 {
						createIds := createIds("US")
						encrpt := encrypt(createParams.UserPassword)
						var Insert = structs.TbUserLogins{
							User_id:       createIds,
							User_name:     createParams.UserName,
							User_uname:    createParams.UserUname,
							User_password: encrpt,
							Role_id:       createParams.UserRole,
							User_plan:     createParams.UserPlant,
							User_status:   1,
							Created_by:    id,
							Updated_by:    id,
						}
						idb.DB.Create(&Insert)
						c.JSON(http.StatusOK, gin.H{
							"status": "ok",
						})
					} else {
						c.JSON(http.StatusBadRequest, gin.H{
							"status":  http.StatusBadRequest,
							"message": "Username Already Exist",
						})
						c.Abort()
					}
				}

			}
		}
	}

}

func (idb *InDB) GetUser(c *gin.Context) {
	token, err := parseBearerToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		fmt.Println("************ERR AUTHORIZED******************")
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, err.Error())
		c.Abort()
	} else {
		decoded := token.Claims.(jwt.MapClaims)
		role := fmt.Sprintf("%v", decoded["role"])
		fmt.Println(role)
		if role != "1" {
			fmt.Println("************ERR IS NOT SU******************")
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
		} else {
			type generateResult struct {
				Numbers     int
				User_id     string
				User_name   string
				User_uname  string
				Plant_label string
				User_status int
				Role_label  string
			}
			var (
				result    gin.H
				users     []structs.TbUserLogins
				Datausers structs.TbUserLogins
				DataPlans structs.TbPlan
				DataRoles structs.TbRoles
				allRes    []generateResult
			)
			queryUser, err := idb.DB.Table("tb_user_logins as users").Select("user_id,user_name,user_uname,plans.plan_label,user_status,roles.role_label").Joins("join tb_plans as plans on users.user_plan = plans.plan_id join tb_roles as roles on users.role_id = roles.role_id").Find(&users).Rows()
			if err != nil {
				fmt.Println("************ERR USERS******************")
				fmt.Println(err)
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  http.StatusBadRequest,
					"message": "Paramater Incomplate",
				})
			} else {
				defer queryUser.Close()
				i := 1
				for queryUser.Next() {
					err := queryUser.Scan(&Datausers.User_id, &Datausers.User_name, &Datausers.User_uname, &DataPlans.Plan_label, &Datausers.User_status, &DataRoles.Role_label)
					if err != nil {
						fmt.Println("************ERR FOR USERS******************")
						fmt.Println(err)
						c.JSON(http.StatusInternalServerError, err)
						c.Abort()
					} else {
						allRes = append(allRes, generateResult{
							Numbers:     i,
							User_id:     Datausers.User_id,
							User_name:   Datausers.User_name,
							User_uname:  Datausers.User_uname,
							Plant_label: DataPlans.Plan_label,
							User_status: Datausers.User_status,
							Role_label:  DataRoles.Role_label,
						})
					}
					i++
				}
				fmt.Println(allRes)
				result = gin.H{
					"result": allRes,
					"count":  len(users),
				}
				c.JSON(http.StatusOK, result)
			}
		}
	}
}
