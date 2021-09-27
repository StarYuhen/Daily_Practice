package MYSQL

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

//查询

// InArchitecture 查询是否有指定数据库架构
func InArchitecture(db *sql.DB, ArchitectureName string) bool {
	SQLString := "show databases"
	res, err := db.Query(SQLString)
	var list []string
	var Name string
	if err != nil {
		logrus.Error("创建数据库报错：", err)
		return false
	}
	//无须关闭数据库，当程序消失时，会自动关闭链接 但只能有一个open
	//defer db.Close()
	for res.Next() {
		res.Scan(&Name)
		list = append(list, Name)
	}

	for _, value := range list {
		if value == ArchitectureName {
			//查询成功会返回名字
			logrus.Info("数据库架构查询:", list)
			return true
		}
	}
	return false
}

// InTableUserConstitute  查询指定账号的所有内容 user_main
func InTableUserConstitute(db *sql.DB, account string) UserMain {
	res, err := db.Query("select * from user_main  where account=?", account)
	if err != nil {
		logrus.Error("数据库查询内容失败user_main:", err)
	}

	var user UserMain

	for res.Next() {
		res.Scan(&user.Account, &user.Password, &user.Authority)
	}

	logrus.Info("查询的是指定内容数据user_main：", user)
	return user
}

// InTableUserUnique 查询是否有该账号
func InTableUserUnique(db *sql.DB, account string) bool {
	var Account string
	_ = db.QueryRow("select account from user_main where account=?", account).Scan(&Account)
	if Account == account {
		return true
	}
	return false
}
