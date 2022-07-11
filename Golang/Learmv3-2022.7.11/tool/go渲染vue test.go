package main

import (
	"html/template"
	"log"
	"net/http"
)

// 文章地址 https://www.jianshu.com/p/2cbf3ed417ef
/*********************************************
 @FuncName:main
 @Author:Jian Junbo
 @Date:2019-11-02 22:12:09
 @Version:1.0
 @Info:This function is 程序的入口，不可以被删除
*********************************************/
func main() {
	http.HandleFunc("/", handler_cofox)
	http.Handle("/htmlpage/", http.StripPrefix("/htmlpage/", http.FileServer(http.Dir("htmlpage"))))
	err := http.ListenAndServe(":1989", nil)
	if err != nil {
		log.Println("ListenAndServe:", err.Error())
	}
}

// News /*********************************************
type News struct {
	Title   string `标题`
	Content string `内容`
	Author  string `作者`
}

func handler_cofox(w http.ResponseWriter, r *http.Request) {
	news := News{Title: "标题", Content: "内容", Author: "作者"}
	t, err := template.ParseFiles("htmlpage/index.html")
	if err != nil {
		log.Println("模板加载失败:", err.Error())
	}
	t.ExecuteTemplate(w, "index.html", news)
}
