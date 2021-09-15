package MYSQL

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

//初始化数据库配置，连接

// MainSQLU 经过处理的连接指定forum_main架构的方法
func MainSQLU(account string, password string) *sql.DB {
	db := MainSql(account, password)
	//一键创建数据库架构forum_main
	CreatArchitectureINit(db, "forum_main")
	//关闭原始数据库
	db.Close()
	//返回指定的数据库内容
	return MainSqlt(account, password)
}

// MainSql 数据库链接初始化地址
func MainSql(account string, password string) *sql.DB {
	db, err := sql.Open("mysql", account+":"+password+"@tcp(localhost:3306)/")
	if err != nil {
		logrus.Error("MYSQL 连接失败 ：", err)
	}
	return db
}

// CreatTableInit 初始化主账号数据库  暂时方案，等后面需要将其改成事务，创建多个表
func CreatTableInit(db *sql.DB) bool {
	SQLString := "create table user_main\n(\n   account text not null,\n    password text not null,\n    authority int null\n);"
	res, err := db.Exec(SQLString)
	if err != nil {
		logrus.Error("数据库表创建失败")
		return false
	}
	logrus.Info("数据库表创建成功", res)
	return true
}

// MainSqlt 链接到指定forum_main架构上
func MainSqlt(account string, password string) *sql.DB {
	db, err := sql.Open("mysql", account+":"+password+"@tcp(localhost:3306)/forum_main")
	if err != nil {
		logrus.Error("数据库链接报错：", err)
	}
	//设置最大连接数 默认为0 也就是没有限制
	//db.SetMaxOpenConns(1000)
	//设置最大空闲连接 每次执行完语句都将连接放入连接池，默认为2  只有open打开才算哦
	db.SetConnMaxIdleTime(10)
	return db
}

//----------------------一键数据库操作类---------------------------------

// CreatArchitectureINit  一键创建指定数据库架构
func CreatArchitectureINit(db *sql.DB, Architecture string) {
	if !InArchitecture(db, Architecture) {
		//当初始化的数据库不存在时创建数据库
		CreatArchitecture(db, Architecture)
	}
}
