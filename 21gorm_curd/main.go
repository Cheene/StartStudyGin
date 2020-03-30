package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math/rand"
	"strconv"
)

//查询功能
type People struct {
	gorm.Model
	ID     int64
	Name   string        `gorm:"default:'xiaowangzi'"`
	Gender sql.NullInt32 // 实现了 Scanner?Valuer 接口
	Age    int64
}

type User struct {
	gorm.Model
	Name   string
	Age    int64
	Active bool
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

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "gormcurd_" + defaultTableName
	}
	//自动迁移地图
	db.AutoMigrate(&People{})
	db.AutoMigrate(&User{})

	//随机初始化
	for i := 0; i < 200; i++ {
		str := "chen_" + strconv.Itoa(rand.Intn(100000))
		db.Create(&People{
			Name: str,
			Age:  rand.Int63n(5455454),
		})
	}
	var people People

	// 根据主键查询最后一条记录
	db.Debug().Last(&people)
	fmt.Printf("Last: %#v\n", &people)

	//查询第一条记录
	db.Debug().First(&people)
	fmt.Printf("First: %#v\n", people)

	//随机获取一条记录
	db.Debug().Take(&people)
	fmt.Printf("Take: %#v\n", people)

	// 查询所有的语句
	db.Debug().Find(&people)
	fmt.Printf("%#v\n", people)
	//按住键查询某一条记录
	db.Debug().Find(&people, 3)
	fmt.Printf("%#v\n", people)
	// WHERE
	var people2 People
	db.Debug().Where("name = ?", "chen_72884").Find(&people2)
	fmt.Printf("%#v\n", people2)

	var people3 []People
	db.Debug().Where("name IN (?)", []string{"chen_81661", "chen_21853"}).Find(&people3)
	fmt.Printf("%#v\n", people3)

	//条件为 Map 或者结构体的类型
	var people4 []People
	db.Debug().Where(map[string]interface{}{"name": "chen_42060", "age": 4182572}).Find(&people4)
	fmt.Printf("%#v\n", people4)
	//按主键返回
	fmt.Println("主键返回")
	var people5 []People
	db.Where([]int64{10, 200, 54}).Find(&people5)
	fmt.Printf("%#v\n", people5)
	fmt.Println("-------------------------------------------------------------------")
	//额外查询选项
	people2.ID = 0
	db.Debug().Set("gorm:query_option", "FOR UPDATE").First(&people2, 10)
	fmt.Printf("%#v\n", people2)

	// Attrs 当记录未找到的时候，将会使用参数初始化
	people5 = nil
	db.Where(People{Name: "chen_52804"}).Attrs(People{Age: 2000}).Find(&people5)
	fmt.Printf("%#v\n", people5)

	//使用 Assign  强制分配
	people5 = nil
	db.Where(People{Name: "chen_52804"}).Assign(People{Age: 2000}).Find(&people5)
	fmt.Printf("%#v\n", people5)

	//使用子查询
	//people5 = nil
	//db.Debug().Where("age > ?",db.Table("gormcurd_peoples").Select("age").Where("name = ?","").QueryExpr()).Find(&people5)
	//fmt.Printf("%#v\n",people5)

	//
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("更新")
	//Save

	db.First(&people)
	people.Name = "nuoning"
	people.Age = 18
	db.Debug().Save(&people)                                                              // 默认修改全部的字段
	db.Debug().Model(&people).Update("name", "yishi")                                     //仅仅更新某个字段
	db.Debug().Model(&people).Updates(map[string]interface{}{"name": "cheen", "age": 20}) //更新多个字段
	//更新选定字段 Select 或者忽略选定字段 Omit
	var people10 People
	db.First(&people10)
	db.Debug().Model(&people10).Select("name").Updates(map[string]interface{}{"name": "cheen", "gender": "1"})
	db.Debug().Model(&people10).Omit("name").Updates(map[string]interface{}{"name": "cheen", "gender": "1"})
	//取消钩子函数
	db.Debug().Model(&people10).UpdateColumn(map[string]interface{}{"name": "cheen", "gender": "1"})
	//数据库某个字段所有的统一 +2
	db.Debug().Model(&People{}).Update("age", gorm.Expr("age + ?", 2))

	//
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("删除")
	// 仅初始化非主键，会全部的软删除
	var p = People{
		Model:  gorm.Model{},
		ID:     1,
		Name:   "",
		Gender: sql.NullInt32{},
		Age:    0,
	}
	db.Debug().Where("id = 2").Delete(&p) //软删除，更新的是 delate_at 字段
	var p1 []People
	db.Unscoped().Find(&p1)
	fmt.Println(p1)
	//物理删除
	//	db.Debug().Unscoped().Delete(&p1)

}
