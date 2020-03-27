package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/**
前提设置了默认值，有两种办法来插入默认值
1 将字段设置为指针类型
2 使用 实现 Scanner/Valuer 的接口
**/

//1 定义模型
//2 把模型与数据库中的表对应起来
//3 CURD
//更改之后需要删除原来的表格
type User struct {
	ID     int64
	Name   *string       `gorm:"default:'xiaowangzi'"`
	Gender sql.NullInt32 // 实现了 Scanner?Valuer 接口
	Age    int64         `gorm:"default:'20''"`
}

func main() {
	//1 连接数据库
	db, err := gorm.Open("mysql", "root:123456@(localhost)/nuoning?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Printf("Err Open db:err %v", err)
		return
	}
	fmt.Println("connect success")
	defer db.Close()
	//添加默认前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "gormcurd_" + defaultTableName
	}

	//2 创建表，自动迁移，把结构体和数据表进行对应
	db.AutoMigrate(&User{})
	//3 创建一个结构体的实例
	u := User{
		Name: new(string),
		Age:  58,
	}
	// 传递 u 也可以，但最好是传递指针
	fmt.Println(db.NewRecord(u)) //判断主键是否为空
	db.Create(&u)
	fmt.Println(db.NewRecord(u)) //判断主键是否为空
	str := "nil"
	u1 := User{
		Name:   &str,
		Gender: sql.NullInt32{1, true},
		Age:    20,
	}
	//添加 Debug() 使用
	db.Debug().Create(&u1)
}
