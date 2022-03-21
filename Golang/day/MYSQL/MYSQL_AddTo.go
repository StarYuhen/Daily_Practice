package MYSQL

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

// 增加

// AddUser 增加用户数据——user_main 表  会自动检测是否有该账号
func AddUser(db *sql.DB, account string, password string, authority int) bool {
	// 检测一下是否有该账号 请求可以被抓包，所以不盲目相信前端
	if InTableUserUnique(db, account) { // 有该账号就直接注册失败
		return false
	}

	res, err := db.Exec("insert into user_main(account, password, authority) VALUE(?,?,?)", account, password, authority)
	if err != nil {
		logrus.Error("增加用户失败user_main:", err)
		return false
	}
	// res.LastInsertId() 是受影响的id  res.RowsAffected() 是受影响的行数
	i, _ := res.RowsAffected()
	return i == 1
}
