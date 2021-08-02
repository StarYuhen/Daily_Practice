package FilesGo

import (
	"encoding/json"
	"github.com/sirupsen/logrus/logrus"
	"io/ioutil"
	"os"
)

type JsonTable struct {
	Db string `json:"db"`
}

func ReadJson() {
	jsonFile, err := os.Open("../SetUp.json")
	if err != nil {
		logrus.Error("读取json配置文件出错:", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var rest map[string]interface{}
	json.Unmarshal([]byte(byteValue), &rest)
	logrus.Info("配置文件内容", rest)
	logrus.Info(rest["db"])
}
