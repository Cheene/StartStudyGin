package dao

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func InitMySql() (err error) {
	DB, err = gorm.Open("mysql", "root:123456@(localhost)/bubble?charset=utf8&parseTime=True&loc=Local")
	return
}

func Close() {
	DB.Close()
}
