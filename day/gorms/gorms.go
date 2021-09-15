package gorms

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//用于对本地数据可，db文件进行操作 sqlite

type GORMS struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	Account   string
	Password  string
	authority int
}

func GormMain() {
	//初始化数据库链接
	db, err := gorm.Open("mysql", "root:abc123456@tcp(localhost:3306)/forum_main/user_main?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		logrus.Error(err)
	}
	gorms := GORMS{Code: "name", Price: 10}
	result := db.Create(&gorms)
	logrus.Info(result.RowsAffected)
	logrus.Info(gorms.ID)

}
