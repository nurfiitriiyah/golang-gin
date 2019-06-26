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

type TbIostock struct {
	Iostock_Dispatch    string
	Iostock_Date        time.Time
	Iostock_type        string
	Iostock_packg       string
	Iostock_in          int
	Iostock_out         int
	Iostock_stok        int
	Iostock_gnt_in      int
	Iostock_que         int
	Iostock_otw         int
	Iostock_osdo        int
	Iostock_update_at   time.Time
	Iostock_last_update time.Time
}

type TbProvid struct {
	Provid_Code      string
	Provid_Name      string
	Provid_Ktgr      string
	Provid_Location  string
	Provid_Cap       int
	Provid_Min       int
	Provid_Agen      string
	Provid_Telp      string
	Provid_Email     string
	Provid_Ip        string
	Provid_Will      string
	Provid_dest      string
	Provid_Out_Pack  string
	Provid_Pasokan   string
	Provid_Seq       int
	Provid_Addr1     string
	Provid_Addr2     string
	Provid_Addr3     string
	Provid_Opeb_Code string
	Provid_Area      string
	Provid_Target    int
	Provid_Ktgr2     string
	Provid_Ktgr3     string
	Provid_Status    string
}

type TbBagcode struct {
	Bagcode_Code  string
	Bagcode_Name  string
	Bagcode_Srt   int
	Bagcode_Ktgr  string
	Bagcode_Bagn  int
	Bagcode_Alias string
	Bagcode_Type  string
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

type UserFirebase struct {
	UserID       string `json:"user_id"`
	Notification string `json:"notification"`
}

/**
HERE DECLARATION FOR EVENT(Combo box or any event in frontend)
**/

type TbProvidEvent struct {
	Provid_Code string
	Provid_Name string
}

type TbBagcodeEvent struct {
	Bagcode_Code string
	Bagcode_Name string
}

/**
HERE DECLARATION FOR REQUEST FROM API
*/

type CreateParams struct {
	Data []string `json:"data"`
}
