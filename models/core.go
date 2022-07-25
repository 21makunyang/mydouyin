package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func init() {
	DB, err = gorm.Open("mysql", "root:mky@(localhost:3306)/mydouyin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	//DB.AutoMigrate(&controller.User{})
	//DB.AutoMigrate(&controller.NameAndPassword{})
}
