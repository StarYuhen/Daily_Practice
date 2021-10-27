package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// jwt只是token的一种使用方式，通常用于鉴权

// JwtPO 自定义jwt参数
type jwtPO struct {
	User     string `json:"user"`
	Password string `json:"password"`
	jwt.StandardClaims
}

const jwtTime = time.Hour * 24

var name = []byte("nmsl")

// SetJwt 生成jwt
func setJwt() string {
	// 因为是本地测试，所以只用打印看看就行
	lie := jwtPO{
		"YUHEN",
		"YYDS",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtTime).Unix(),
			Issuer:    "Yuhen",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, lie)
	tokens, _ := token.SignedString(name)
	fmt.Println("加密过后的jwt", tokens)
	return tokens
}

// 解析jwt
func jwtread(jwts string) (*jwtPO, error) {
	// 解析原始内容
	token, err := jwt.ParseWithClaims(jwts, &jwtPO{}, func(token *jwt.Token) (interface{}, error) {
		return name, nil
	})
	if err != nil {
		return nil, err
	}
	// 将解析后的内容返回
	if clime, ok := token.Claims.(*jwtPO); ok && token.Valid {
		return clime, nil
	}

	return nil, errors.New("宝啊，可不兴这样的")
}

func main() {
	token := setJwt()
	c, _ := jwtread(token)
	fmt.Println("解析jwt后的内容", c)
}
