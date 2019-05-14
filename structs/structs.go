package structs

import "github.com/jinzhu/gorm"

/**
every struct name must use capital in first character
**/
type Person struct {
	gorm.Model
	First_Name string
	Last_Name string
}