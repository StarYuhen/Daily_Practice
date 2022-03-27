package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// 用于对本地数据可，db文件进行操作 sqlite

type GORMS struct {
	gorm.Model
	Code  string
	Price uint
}

// User 使用改结构体创建数据表，创建的表名是蛇形，这个创建的名称叫users
type User struct {
	ID        string // id会自动生成索引
	Account   string `gorm:"primaryKey"` // `gorm:"primarykey"` 会将值变为主键
	Password  string
	Authority int
}

// USER 使用标签自动创建表
type USER struct {
	ID string `gorm:"size:255"`
}

func main() {
	// 初始化数据库链接
	db, err := gorm.Open("mysql", "root:abc123456@tcp(localhost:3306)/forum_main?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		logrus.Error(err)
	}
	// autoMigrate 自动迁移数据库 也看作自动生成数据表，他并不会把上一个表的数据迁移过来
	// db.AutoMigrate(&User{})
	db.AutoMigrate(&USER{})
	list1 := User{Account: "wdnmd", Password: "sdfsd", Authority: 100}
	// 创建一个记录
	result := db.Create(&list1)
	/*
		list1.ID             // 返回插入数据的主键
		result.Error        // 返回 error
		result.RowsAffected // 返回插入记录的条数
	*/

	logrus.Info(result.Error)
	logrus.Info(result.RowsAffected)
	// 设置表后缀
	db.Set("gorm:gorm_juan", "ENGINE=InnoDB").AutoMigrate(&User{})

	logrus.Info(db.RowsAffected)

}
