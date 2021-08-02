package MYSQL

import (
	"database/sql"
	"github.com/sirupsen/logrus/logrus"
)

//更新数据库内容

// UpdateTableUser  修改，账号，密码，权限
func UpdateTableUser(db *sql.DB, account string, password string, authority int) bool {
	res, err := db.Exec("update user_main  set account=? where password=? and authority=?", account, password, authority)
	if err != nil {
		logrus.Error("更新数据库账号密码错误user_main：", err)
		return false
	}
	//受到影响的行数
	i, _ := res.RowsAffected()
	return i == 1
}
