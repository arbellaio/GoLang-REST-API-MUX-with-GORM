package dbContext


import (
	"fmt"
	model "goLangRestApiMuxGorm/Models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
const DatabaseUrl = "root:faizan123@tcp(127.0.0.1)/godb?charset=utf8mb4&parseTime=True&loc=Local"

func InitialMigration()  {
	DB, err = gorm.Open(mysql.Open(DatabaseUrl), &gorm.Config{})
	if err != nil {
		 fmt.Println(err.Error())
		 panic("Cannot connect to DB")
	}

	DB.AutoMigrate(&model.User{})
}