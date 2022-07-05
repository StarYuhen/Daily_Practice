package main

import (
	"fmt"
	"github.com/gogf/gf/util/grand"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type PhoneAccountStruct struct {
	ID              uint   `gorm:"primaryKey"`
	Guid            string `json:"Guid,omitempty"`
	Phone           string `json:"Phone,omitempty"`
	Mid             string `json:"Mid,omitempty"` // 根据MID 和 GUid 这个两个字段更新 ==mid and ==guid
	Sent            bool   `json:"Send,omitempty"`
	Read            bool   `json:"Read,omitempty"`
	Error           string `gorm:"column:error"`
	SendImg         string `json:"SendImg,omitempty" gorm:"column:send_img"`   // 这个是头像
	SendName        string `json:"SendName,omitempty" gorm:"column:send_name"` // 这个是用户昵称
	TaskMid         string `gorm:"column:task_mid"`                            // 属于哪个任务用掉的账号
	ReadMessageBool bool   `gorm:"column:read_message_bool"`                   // 消息是否已读
	CreateTime      string `gorm:"column:create_time"`                         // 发送时间
	SendID          string `gorm:"column:send_id"`                             // 更新的对方id
	TaskTokenUID    string `gorm:"column:task_token_uid"`
}

// GormDB defaa84a8fa6bcec
var GormDB = SqlInit("root", "abc123456", "golandfile")
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func SqlInit(User, Password, Architecture string) *gorm.DB {
	sqlInit := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?utf8mb4&parseTime=True&loc=Local", User, Password, Architecture)
	SQL, err := gorm.Open(mysql.Open(sqlInit), &gorm.Config{})
	if err != nil {
		logrus.Error("数据库链接报错：", err)
		return nil
	}

	SQL.InstanceSet("gorm:table_options", "ENGINE=InnoDB")

	// 利用database/sql设置数据库连接池
	sql, err := SQL.DB()

	// 设置最大连接数 默认为0 也就是没有限制
	sql.SetMaxOpenConns(0)
	// 设置最大空闲连接 每次执行完语句都将连接放入连接池，默认为2
	sql.SetConnMaxIdleTime(100000)
	logrus.Info("已初始化mysql链接")
	return SQL
}

type MysqlUpdateTaskMid struct {
	Mid         string
	SendUIDTim  string
	PicturePath string
	DisplayName string
	ID          string
	MId         string
	ReceiverId  string
}

func main() {

	var p []PhoneAccountStruct
	var m = []MysqlUpdateTaskMid{{ReceiverId: "e24b4c4e-f24e-4373-910d-65621c0b297c", Mid: "u7ccb4c74c0dc58358ac38e652befbc52"}}

	var TaskMid = make(map[string]MysqlUpdateTaskMid, 0)

	var mid []string
	// 通过遍历写入map
	for _, v := range m {
		TaskMid[v.Mid] = v
		mid = append(mid, v.Mid)
	}

	// mids := []string{"u7ccb4c74c0dc58358ac38e652befbc52", "uee75153a8bd4584c7023e380315628da"}

	// "send_img":       v.PicturePath,
	// 	"send_name":      v.DisplayName,
	// 	"task_mid":       v.SendUIDTim,
	// 	"create_time":    CreateTime,
	// 	"send_id":        v.ID,
	// 	"task_token_uid": v.MId,

	// 使用临时表进行批量更新
	// var p []PhoneAccountStruct

	GormDB.Transaction(func(tx *gorm.DB) error {
		// 获取当前表的所有数据
		tx.Table("phone_account").Where("mid in ?", mid).Having("guid=?", m[0].ReceiverId).Find(&p)
		// 创建临时表
		uuid := randStr(10)
		// 拼接sql语句
		sqlExec := fmt.Sprintf("update phone_account p inner join (select id,send_img,send_name,task_mid,create_time,send_id,task_token_uid from %s) ps on p.id=ps.id set p.send_img=ps.send_img, p.send_name=ps.send_name, p.task_mid=ps.task_mid, p.create_time=ps.create_time, p.send_id=ps.send_id,\n    p.task_token_uid=ps.task_token_uid", uuid)
		sqlDrop := fmt.Sprintf("drop table %s;", uuid)
		// 将结果for循环写入数组类型，方便创建
		for i, v := range p {
			if value, ok := TaskMid[v.Mid]; ok {
				p[i].SendImg = value.PicturePath
				p[i].SendName = value.DisplayName
				p[i].TaskMid = value.SendUIDTim
				p[i].CreateTime = ""
				p[i].SendID = value.ID
				p[i].TaskTokenUID = value.MId
			}
		}

		tx.Table(uuid).AutoMigrate(&PhoneAccountStruct{})
		// 批量创建读取的数据到临时表
		tx.Table(uuid).Create(&p)

		tx.Table(uuid).Exec(sqlExec)

		tx.Table(uuid).Exec(sqlDrop)
		// 然后执行sql语句互相更新
		return nil
	})
}

func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[grand.Intn(len(letters))]
	}
	return string(b)
}
