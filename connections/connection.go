package connections

import (
	"day2/structs"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//bikin variable buat makai function gorm
var (
	DB  *gorm.DB
	Err error
)

func Connect() {
	//connect db ke sql
	DB, Err = gorm.Open("mysql", "root:@/day2?charset=utf8&parseTime=True")

	if Err != nil {
		log.Println("Connection failed", Err)
	} else {
		log.Println("Server up and running")
	}
	//migrate user & risk_profile
	DB.AutoMigrate(&structs.User{})
	DB.AutoMigrate(&structs.Risk_profile{})
}
