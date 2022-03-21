package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

// 测试aes加密

// 密钥长度为16--对于aes-128
var commonIV = []byte("12345678abcdefgh")

func main() {
	// 需要去加密的字符串
	plaintext := []byte("My name is Astaxie")
	// 如果传入加密串的话，plaint 就是传入的字符串
	if len(os.Args) > 1 {
		plaintext = []byte(os.Args[1])
	}

	// aes的加密字符串
	key_text := "astaxie12798akljzmknm.ahkjkljl;k"

	fmt.Println(len(key_text))

	// 创建加密算法 aes
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)
		os.Exit(-1)
	}

	// 加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("%s=>%x\n", plaintext, ciphertext)

	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)
}
