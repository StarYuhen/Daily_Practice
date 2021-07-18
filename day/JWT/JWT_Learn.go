package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

//利用gin框架并使用jwt保存信息，jwt主要用于鉴权

/*
  定义JWT参数
  jwt.StandardClaims 包含了jwt包的官方字段
  我们需要增加自定义的字段 如Username ，Userpassword。
  之所以是json形式，是因为jwt的概念
  jwt是一个用于鉴权的方式，区别于传统的session ，他会自动生成一个json格式信息储存在浏览器用于请求头自动携带。 json web token
  ps:json的自定义参数结构开头必须大写，表示公共参数 public ，如：Username
  jwt详解 https://www.ruanyifeng.com/blog/2018/07/json_web_token-tutorial.html
*/
type jwtJson struct {
	Username     string `json:"username"`
	Userpassword string `json:"userpassword"`
	jwt.StandardClaims
}

// JwtTokenTime 定义jwt过期时间 ，设置为24小时
const jwtTokenTime = time.Hour * 24

// JwtSecret 定义secret
var jwtSecret = []byte("测试用的jwt")

// JwtSet 生成jwt信息
func jwtSet(username string, userpassword string) (string, error) {
	//填充字段
	Set := jwtJson{
		username,
		userpassword,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtTokenTime).Unix(), //设置时间
			Issuer:    "yuhen",                             //设置签发人
		},
	}

	//创建指定签名方法的签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodES256, Set)
	//使用指定签名后的jwt字符串
	fmt.Println("jwt的内容:", token)
	return token.SignedString(jwtSecret)
}

// JwtRead 用于解析jwt
func jwtRead(tokenjson string) (*jwtJson, error) {
	//解析token  自己鼠标悬浮jwt.ParseWithClaims来看
	token, err := jwt.ParseWithClaims(tokenjson, &jwtJson{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if climas, ok := token.Claims.(*jwtJson); ok && token.Valid { //校验token，进行比对
		return climas, nil
	}

	return nil, errors.New("宝啊，可不兴这样的")
}

// JwtFrome gin使用jwt 用户成功登录后会自动在请求头携带jwt的
func JwtFrome(context *gin.Context) {
	/*
		//获取请求接口表单的用户密码
		jwtJsonFrome := make(map[string]interface{})
		err := context.BindJSON(jwtJsonFrome)
		if err != nil {
			return
		}
		username := jwtJsonFrome["user"].(string)
		userpass := jwtJsonFrome["userpass"].(string)

	*/
	//当账号密码与内置的一样时
	//	if username == "admin" && userpass == "admin" {
	if true {
		//tokenstring, _ := JwtSet(username, userpass)
		tokenstring, _ := jwtSet("admin", "admin")
		data := map[string]interface{}{
			"code": 1000,
			"msg":  "Achenes",
			"data": gin.H{"token": tokenstring},
		}
		fmt.Println("jwt的内容:", tokenstring)
		context.JSON(http.StatusOK, data)
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 1002,
		"msg":  "鉴权失败",
	})
	return
}

// JwtMild 校验jwt的中间件
func JwtMild() func(context *gin.Context) {
	return func(context *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		autoHead := context.Request.Header.Get("Authorization")
		if autoHead == "" {
			context.JSON(http.StatusOK, gin.H{
				"code": 1003,
				"msg":  "请求头auto为空",
			})
			context.Abort()
			return
		}

		//按空格分割
		past := strings.SplitN(autoHead, "", 2)
		if !(len(past) == 2 && past[0] == "Bearer") {
			context.JSON(http.StatusOK, gin.H{
				"code": 1004,
				"msg":  "请求头中auto有错",
			})
			context.Abort()
			return
		}

		list, err := jwtRead(past[1])
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"code": 1005,
				"msg":  "无效的Token",
			})
			context.Abort()
			return
		}

		context.Set(list.Username, list.Userpassword)
		context.Next()

	}
}

func main() {
	run := gin.Default()
	run.POST("/jwt", JwtFrome)
	//为这个路径请求加载中间件
	//	sv:=run.Group("/")
	//sv.Use(JwtMild())
	//设置端口
	run.Run(":8000")
}
