package structs

import (
	"github.com/dgrijalva/jwt-go"
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
	jwt.StandardClaims
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

type TbOutstanding struct {
	Outstanding_transporter string
	Outstanding_kategori    string
	Outstanding_dispatcher  string
	Outstanding_area        string
	Outstanding_quantity    int
	Outstanding_update      time.Time
	Outstanding_location    string
	Outstanding_late        int
	Outstanding_package     string
	Outstanding_retail      int
}

type TbDelivery struct {
	Delivery_Date        time.Time
	Delivery_Transporter string
	Delivery_Dispatcher  string
	Delivery_Area        string
	Delivery_Quantity    int
	Delivery_SalesType   string
	Delivery_Update      time.Time
	Delivery_Location    string
	Delivery_Retail      string
	Delivery_Late        int
	Delivery_Package     string
	Delivery_Zona        string
	Delivery_Dest        string
}

type TbLog struct {
	Log_Id     string
	Log_Detail string
	Log_Resp   string
	Created_at time.Time
}

type TbRetail struct {
	Retail_ID    int
	Retail_label string
}

type DetailData struct {
	Data string `json:"data"`
}
