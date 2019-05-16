package structs

import (
	"github.com/jinzhu/gorm"
	"time"
)

/**
every struct name must use capital in first character
**/
type Person struct {
	gorm.Model
	First_Name string
	Last_Name  string
}

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/**
Migration Tabel User
**/
type TbUserLogins struct {
	User_id       string `gorm:"unique;not null"`
	User_name     string
	User_uname    string
	User_password string
	Role_id       int
	User_plan     string
	User_status   int
	Created_at    time.Time
	Created_by    string
	Updated_at    time.Time
	Updated_by    string
}
