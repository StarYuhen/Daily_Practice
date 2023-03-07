package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize"
	"strconv"
)

func main() {
	excel, err := excelize.OpenFile("E:\\CodeDemo\\Daily Practice\\Golang\\Learnv2\\test.xlsx")
	logrus.Error(err)

	for i := 2; i <= 31; i++ {
		score := 0
		number := strconv.Itoa(i)
		cells := []string{"Z" + number, "AB" + number, "AC" + number, "AD" + number, "AE" + number, "AF" + number, "AG" + number, "AI" + number}
		for _, cell := range cells {
			value, err := excel.GetCellValue("Sheet1", cell)
			if err != nil {
				// 处理错误
				logrus.Error("读取单元格错误")
				return
			}
			num, _ := strconv.Atoi(value)
			if num >= 95 {
				score = score + 4
			} else if num >= 90 {
				score = score + 3
			} else if num >= 80 {
				score = score + 2
			}

		}

		// 设置单元格的值
		excel.SetCellInt("Sheet1", "AK"+number, score)
	}
	if err := excel.SaveAs("E:\\CodeDemo\\Daily Practice\\Golang\\Learnv2\\Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
