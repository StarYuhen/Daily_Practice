package main

import (
	"fmt"
	"net/url"
	"strconv"
)

/*
biz_content=%257B%2522total_amount%2522%253A1%252C%2522trade_status%2522%253A%2522SUCCESS%2522%252C%2522extra%2522%253A%2522f5fc8df561e348378c4b79ef3b567aa0%2522%252C%2522trade_pay_time%2522%253A%25222021-10-14+16%253A12%253A29%2522%252C%2522title%2522%253A%2522%25E9%25A6%2596%25E6%25AC%25A1%25E5%2585%2585%25E5%2580%25BC6%25E5%2585%2583%2522%252C%2522order_id%2522%253A%2522250532372034848432%2522%252C%2522biz_order_id%2522%253A%25221_6937_6712_21_16341991448997%2522%257D&nonce_str=2e78f8bafbf94a8494cae8121a2b6410&sign=5af66ffa0252008d46c5e4c36ec40690&app_id=9b7d68729b934e448699afc9a978501d&sign_type=MD5&version=1.0&client_id=60e73b8892ec49ef8b06685eabceb6e2&timestamp=1634199155366xx
*/
func main() {

	c := "biz_content=%257B%2522total_amount%2522%253A1%252C%2522trade_status%2522%253A%2522SUCCESS%2522%252C%2522extra%2522%253A%2522f5fc8df561e348378c4b79ef3b567aa0%2522%252C%2522trade_pay_time%2522%253A%25222021-10-14+15%253A49%253A37%2522%252C%2522title%2522%253A%2522%25E9%25A6%2596%25E6%25AC%25A1%25E5%2585%2585%25E5%2580%25BC6%25E5%2585%2583%2522%252C%2522order_id%2522%253A%2522250529495266581804%2522%252C%2522biz_order_id%2522%253A%25221_6937_6712_21_16341977729852%2522%257D&nonce_str=f276ea4e35c04da7917e825122f7ad7e&sign=bfcfa9436c764ffb1c018dee82b030e8&app_id=9b7d68729b934e448699afc9a978501d&sign_type=MD5&version=1.0&client_id=60e73b8892ec49ef8b06685eabceb6e2&timestamp=1634198150343"

	// i:=[]byte(c)
	str, _ := url.ParseQuery(c)
	str["biz_content"][0], _ = url.QueryUnescape(str["biz_content"][0])
	// i:=[]byte(l)
	/*
		str:="biz_content="+l["biz_content"][0]+"&nonce_str="+l["nonce_str"][0]+"&sign="+l["sign"][0]
		fmt.Println(l)
		fmt.Println(str)

	*/

	//	fmt.Println(url.QueryUnescape(c))
	// fmt.Println(url.QueryUnescape(c))
	// fmt.Println(url.QueryEscape(c))
	// 直接拼接字符串

	strbody := "biz_content=" + str["biz_content"][0] + "&nonce_str=" + str["nonce_str"][0] + "&sign=" + str["sign"][0] + "&app_id=" + str["app_id"][0] + "&sign_type=" + str["sign_type"][0] + "&version=" + str["version"][0] + "&client_id=" + str["client_id"][0] + "&timestamp=" + str["timestamp"][0]
	ls := []byte(strbody)
	fmt.Println(str)
	fmt.Println(string(ls))
	fmt.Println(strconv.ParseInt(str["timestamp"][0], 10, 64))

}

type Bo struct {
	Biz_content string
	Nonce_str   string
	Sign        string
	App_id      string
	Sign_type   string
	Version     string
	Client_id   string
	Timestamp   string
}
