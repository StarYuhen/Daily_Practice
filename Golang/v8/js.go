package main

import (
	"fmt"
	_ "github.com/dop251/goja"
	_ "github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
	v8 "rogchap.com/v8go"
)

// "github.com/dop251/goja" 使用此项目的虚拟机（不太满足要求）
// func maingoja() {
// 	// 启用js虚拟机
// 	vm := goja.New()
// 	// 运行js代码
// 	v, err := vm.RunString("2 + 2")
// 	if err != nil {
// 		panic(any(err))
// 	}
// 	if num := v.Export().(int64); num != 4 {
// 		panic(any(num))
// 	}
// 	logrus.Info(v)
// }

// "github.com/robertkrimen/otto" 不满意其功能
// func mainotto() {
// 	vm := otto.New()
// 	run, err := vm.Run(`
// const request = require('request');
// 	  let myJson = {
// 	        "siteUrl": "https://www.yuhenm.com",
// 	        "urlList": [
// 	            url
// 	        ]
// 	    };
// 	    request({
// 	        url: 'https://ssl.bing.com/webmaster/api.svc/json/SubmitUrlbatch?apikey=' + '25f65107c2fc44ee9663191662e1cd90', /* xxx需替换为你的key */
// 	        method: "POST",
// 	        json: true,   // <--Very important!!!  允许跨域
// 	        body: myJson
// 	    }, function (error, response, body) {
// 	        console.log('提交必应结果',body);
// 	    });
// 	console.log(10)
// 	`)
// 	if err != nil {
// 		logrus.Info(err)
// 	}
// 	logrus.Info("run:", run)
// }

/*
D:/minge/bin/../lib/gcc/x86_64-w64-mingw32/8.1.0/../../../../x86_64-w64-mingw32/bin/ld.exe: cannot find -lv8
collect2.exe: error: ld returned 1 exit status 报错
采用安装msysy2实现win调用v8
虽然没有为 Windows 包含预构建的静态 V8 库，但 MSYS2 提供了一个包，其中包含一个可以工作的动态链接的 V8 库。

要进行此设置：

安装 MSYS2 ( https://www.msys2.org/ )
将 Mingw-w64 bin 添加到您的 PATH 环境变量（C:\msys64\mingw64\bin默认情况下）
打开 MSYS2 MSYS 并执行 pacman -S mingw-w64-x86_64-toolchain mingw-w64-x86_64-v8
这将允许构建依赖于v8go.snapshot_blob.binmain.go
V8 在 Windows 上需要 64 位，因此不适用于 32 位系统。
*/
func main() {
	ctx := v8.NewContext() // creates a new V8 context with a new Isolate aka VM
	con := v8.NewContext()
	ctx.RunScript("const add = (a, b) => a + b", "math.js") // executes a script on the global context
	ctx.RunScript("const result = add(3, 4)", "main.js")    // any functions previously added to the context can be called
	// 试用nodejs
	con.RunScript("const fs = require('fs'); \n fs.mkdirSync(`Rabbit`);", "ccs.js")
	val, _ := ctx.RunScript("result", "value.js") // return a value in JavaScript back to Go
	vals, _ := con.RunScript("fs", "tl.js")
	fmt.Printf("addition result: %s,%s", val, vals)
}
