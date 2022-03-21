package main

import (
	"encoding/json"
	fmt "fmt"
	"io/ioutil"
	http "net/http"
	"strings"
	"time"
)

// 注册信息
type regInfo struct {
	Phone        string
	Cc           string
	OsName       string
	OsVersion    string
	DeviceName   string
	Manufacturer string
}

func newRegInfo(Phone, Cc, OsName, OsVersion, DeviceName, Manufacturer string) *regInfo {
	return &regInfo{
		Phone:        Phone,
		Cc:           Cc,
		OsName:       OsName,
		OsVersion:    OsVersion,
		DeviceName:   DeviceName,
		Manufacturer: Manufacturer,
	}
}

// 获取验证码
func getVerifyCode(url string) string {
	var i = 7
	var code string
	for i > 0 {
		result, err := http.Get(url)
		if err != nil {
			fmt.Println("http请求失败 ", err)
		}
		resultData, err := ioutil.ReadAll(result.Body)
		if err != nil {
			fmt.Println("ioutil.ReadAll fail ", err)
		}
		codeResult := string(resultData)
		fmt.Println(codeResult)
		if codeResult == "STATUS_WAIT_CODE" || codeResult == "STATUS_CANCEL" || codeResult == "" {
			fmt.Println("正在等待短信,10秒后请求")
			time.Sleep(time.Second * 10)
		}
		index := strings.Index(codeResult, ":")
		code = codeResult[index+1:]

		fmt.Printf("code: %s, len: %d ", code, len([]rune(code)))
		if len([]rune(code)) == 6 {
			break
		} else {
			code = ""
		}
		i--
	}

	return code
}

// 服务器生成实例并返回uuid
func serverNumberExists(serverUrlExists string, data []byte) string {
	// json
	contentType := "application/json"
	resp, err := http.Post(serverUrlExists, contentType, strings.NewReader(string(data)))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer resp.Body.Close()
	phoneBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(phoneBody))
	var phoneDataJson = make(map[string]interface{}, 1)
	if err := json.Unmarshal([]byte(phoneBody), &phoneDataJson); err != nil {
		fmt.Println(err)
		return ""
	}
	if phoneDataJson["Code"].(float64) != 0 {
		fmt.Println("phoneDataJson[\"Code\"] ！= 0 ")
		return ""
	}

	return phoneDataJson["Msg"].(string)
}

// 获取号码
func getPhone(url string) string {
	data, err := http.Get(url)
	if err != nil {
		fmt.Println("获取号码失败", err)
		return ""
	}
	dataBody, err := ioutil.ReadAll(data.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll(data.Body) ", err)
		return ""
	}
	result := string(dataBody)
	fmt.Println(result)
	if result == "NO_NUMBERS" || result == "NO_BALANCE" {
		fmt.Println("没有号码或者余额结束")
		return ""
	}
	return result
}

// 发送验证码
func sentVerifyCode(url string) ([]byte, error) {
	get, err := http.Get(url)
	if err != nil {
		fmt.Println("get request failed, err:", err)
		return nil, err
	}
	verifyCodeData, err := ioutil.ReadAll(get.Body)
	if err != nil {
		fmt.Println("get data failed, err:", err)
		return nil, err
	}
	return verifyCodeData, nil
}

// 向服务器发送验证码
func reg(url, code string) ([]byte, error) {
	// json
	contentType := "application/json"
	data := "{\"Code\":\"" + code + "\"}"
	fmt.Println(data)
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("post failed, err: ", err)
		return nil, err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed, err:", err)
		return nil, err
	}
	return result, nil
}

func main() {
	const serverUrlVerifyCode = "http://127.0.0.1:8011/ws/reg/verifyCode/"
	const serverUrlExists = "http://127.0.0.1:8011/ws/reg/exists"
	const serverUrlReg = "http://127.0.0.1:8011/ws/reg/reg/"
	const PhoneBaseurl = "https://sms-activate.ru/stubs/handler_api.php"
	const apikey = "96df4529261c6Abd84dA4e3603cA9d00"
	// whatsapp
	const service = "wa"
	// 印尼
	const country = "6"
	// 获取手机号和id
	getPhoneNumberUrl := PhoneBaseurl + "?api_key=" + apikey + "&action=getNumber&service=" + service + "&country=" + country
	phoneResult := getPhone(getPhoneNumberUrl)
	if phoneResult == "" {
		return
	}
	// 数据截取
	phoneNumberIndex := strings.LastIndex(phoneResult, ":")
	phoneIdIndex := strings.Index(phoneResult, ":")
	id := phoneResult[phoneIdIndex+1 : phoneNumberIndex]
	phoneNumber := phoneResult[phoneNumberIndex+3:]

	regInfo := newRegInfo(phoneNumber, "62", "Android", "7.1.2",
		"fortunave3222", "samsung")
	data, err := json.Marshal(regInfo)
	if err != nil {
		fmt.Println("json marshal failed, err:", err)
		return
	}

	uuid := serverNumberExists(serverUrlExists, data)
	sentVerifyUrl := serverUrlVerifyCode + uuid
	fmt.Println(id)
	fmt.Println(phoneNumber)
	fmt.Println(sentVerifyUrl)
	var sentVerifyCodeRes = make(map[string]interface{}, 1)

	var serverResultJson = make(map[string]interface{}, 1)
	for i := 0; i < 4; i++ {
		// 发送验证码，第一次有可能返回-4 所以重复发送，一直到返回0为止
		for {
			codeData, err := sentVerifyCode(sentVerifyUrl)
			if err != nil {
				fmt.Println("请求失败，服务未打开")
				return
			}
			if err := json.Unmarshal(codeData, &sentVerifyCodeRes); err != nil {
				fmt.Println("json unmarshal failed, err: ", err)
			}
			fmt.Println(string(codeData))
			if sentVerifyCodeRes["Code"].(float64) == 0 {
				break
			}
		}

		getVerifyCodeUrl := PhoneBaseurl + "?api_key=" + apikey + "&action=getStatus&id=" + id
		fmt.Println(getVerifyCodeUrl)
		// 获取验证码
		verifyCode := getVerifyCode(getVerifyCodeUrl)
		regUrl := serverUrlReg + uuid
		fmt.Println(verifyCode)
		fmt.Println(regUrl)
		if verifyCode != "" {
			for {
				// 向服务器提交验证码
				serverResultB, err := reg(regUrl, verifyCode)
				if err != nil {
					fmt.Println("请求失败")
					return
				}
				// 将服务器返回的string反序列化，看code 结果
				if err := json.Unmarshal(serverResultB, &serverResultJson); err != nil {
					fmt.Println("json unmarshal failed, err:", err)
					return
				}
				// fmt.Println(string(serverResultB))
				fmt.Println(serverResultJson["Code"])
				if serverResultJson["Code"].(float64) == 0 {
					fmt.Println(string(serverResultB))
					break
				} else {
					fmt.Println("错误码 code:", serverResultJson["Code"])
					fmt.Println(string(serverResultB))
				}
			}

			break
		}
	}

}
