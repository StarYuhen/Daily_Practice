package Go

import (
	"../day/FilesGo"
	"testing"
)

func TestJson(t *testing.T) {
	//这里需要注意的准则时，前面引用部分不能用../ ，倘若要引用其他文件请生成mod文件然后引用这个地址的文件，不要用相对路径引用
	FilesGo.ReadJson()
}
