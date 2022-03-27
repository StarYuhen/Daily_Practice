package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type APIResp struct {
	Code int
	Data map[string]interface{}
	Msg  string
}

func main() {
	str := "{\r\n    \"Socks5\": \"%s\",\r\n    \"EdgeRouting\": \"CA0IBQ==\",\r\n    \"AuthBody\": {\r\n        \"username\": 6283874726273,\r\n        \"passive\": false,\r\n        \"user_agent\": {\r\n            \"platform\": 0,\r\n            \"app_version\": {\r\n                \"primary\": 2,\r\n                \"secondary\": 21,\r\n                \"tertiary\": 12,\r\n                \"quaternary\": 22\r\n            },\r\n            \"mcc\": \"310\",\r\n            \"mnc\": \"070\",\r\n            \"os_version\": \"7.1.2\",\r\n            \"manufacturer\": \"samsung\",\r\n            \"device\": \"fortunave32\",\r\n            \"os_build_number\": \"google-user 7.1.2 20171130.276299 release-keys\",\r\n            \"phone_id\": \"e31e46e0-f701-42ff-a868-a81c5785c2a9\",\r\n            \"locale_language_iso_639_1\": \"en\",\r\n            \"local_country_iso_3166_1_alpha_2\": \"US\",\r\n            \"Brand\": \"msm8998\"\r\n        },\r\n        \"push_name\": \"sjeyg\",\r\n        \"short_connect\": false,\r\n        \"connect_type\": 1,\r\n        \"Tag23\": 1\r\n    },\r\n    \"AuthHexData\": \"\",\r\n    \"StaticPriKey\": \"SMf8lZ5YKhAi/CAFM6GT/KfUGY8Mqfrogse8sKK7SEg=\",\r\n    \"StaticPubKey\": \"/Z4o4acNy7cQjGoN6hM2ZTOiRt1hPLMYEm8y8Prs030=\"\r\n}"
	str = fmt.Sprintf(str, "wsdfdsf")
	logrus.Info(str)
}
