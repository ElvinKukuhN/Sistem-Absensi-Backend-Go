package Database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	const mySql = "root@tcp(127.0.0.1:3306)/sistem_absensi?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := mySql
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = d
	fmt.Println("Connection Opened to Database")
}
