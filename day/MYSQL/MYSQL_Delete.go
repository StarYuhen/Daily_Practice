package MYSQL

import (
	"database/sql"
	"github.com/sirupsen/logrus/logrus"
)

//删除数据库内容

func DeleteTableUser(db *sql.DB, account string) bool {
	res, err := db.Exec("delete from user_main where account=?", account)
	if err != nil {
		logrus.Error("删除账号失败user_main:", err)
		return false
	}

	i, _ := res.RowsAffected()
	return i == 1
}
