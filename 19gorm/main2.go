package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

//go.Model 可以包括 正常定义的结构体;基本的go类型或他们的指针;
//或者 sql.Scanner 以及 driver.Valuer 接口
type UserInfo struct {
	gorm.Model //gorm.Model
	ID         uint
	Name       string
	Gender     string
	Hobby      string
}

//模型的示例
type BuildInfo struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`
	MemberNumber *string `gorm:"unique;not null"`
	Num          int     `gorm:"AUTO_INCREMENT"`
	Address      string  `gorm:"index:addr"` //创建 名为 addr 的索引
	IgnoreMe     int     `gorm:"-"`          //忽略此字段
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
		return "chen_" + defaultTableName
	}

	//创建表，自动迁移，把结构体和数据表进行对应
	db.AutoMigrate(&UserInfo{})
	db.AutoMigrate(&BuildInfo{})
	u1 := UserInfo{
		ID:     2,
		Name:   "Chenene",
		Gender: "male",
		Hobby:  "xxx",
	}
	db.Create(&u1)

	u2 := UserInfo{
		ID:     1,
		Name:   "admin",
		Gender: "male",
		Hobby:  "xxx",
	}
	db.Create(&u2)

	//查询
	var u UserInfo
	db.First(&u)
	fmt.Printf("u:%#v\n", &u)
	//更新
	db.Model(&u).Update("hobby", "pingpangqiu")
	//删除
	db.Delete(&u)

	fmt.Println("connect success")
}
