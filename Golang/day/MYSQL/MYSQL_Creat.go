package MYSQL

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

// CreatArchitecture 创建数据库架构
func CreatArchitecture(db *sql.DB, ArchitectureName string) bool {
	SQLString := "create schema " + ArchitectureName + ";"
	res, err := db.Exec(SQLString)
	if err != nil {
		logrus.Error("创建数据库报错：", err)
		return false
	}
	i, _ := res.RowsAffected()
	return i == 1
}
