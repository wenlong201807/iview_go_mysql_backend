package main

import (
	"commons"
	"content"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"item"
	"item/cat"
	"item/param"
	"item/paramitem"
	"net/http"
	"user"
)

//显示登录页面
func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/login.html")
	t.Execute(w, nil)

}

// restful显示页面
func showPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t, _ := template.ParseFiles("view/" + vars["page"] + ".html")
	t.Execute(w, nil)
}

func main() {

	commons.Router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	commons.Router.HandleFunc("/", welcome)
	// 满足/page/{page} 格式的处理// 前后端不分离项目需要的
	commons.Router.HandleFunc("/page/{page}", showPage)
	user.UserHandler()           // 调用登录的接口
	item.ItemHandler()           // 调用商品信息的接口 // 查询商品
	cat.ItemCatHandler()         // 调用商品信息，添加商品，，，选择类目的接口
	content.ContentHandler()     // 调用内容信息
	param.ParamHandler()         // 规格参数分页显示
	paramitem.ParamItemHandler() // 商品规格参数
	http.ListenAndServe("192.168.43.148:80", commons.Router) // 公司
	//http.ListenAndServe("192.168.31.30:80", commons.Router)
	//http.ListenAndServe(":80", commons.Router)
	fmt.Println("8888")
}

/*
func main() {
	fmt.Println("开始执行接口数据了。。。。")
	//s := http.Server{Addr: ":8089"}
	s := http.Server{Addr: "http://192.168.43.148:8089"}
	//http.Handle("/static/",http.StripPrefix("/static",http.FileServer(http.Dir("static"))))
	//http.HandleFunc("/",welcome)
	//// 调用所有user模块的handler
	user.UserHandler() // 调用登录的接口
	//	item.ItemHandler()  // 调用商品信息的接口 // 查询商品
	//	cat.ItemCatHandler() // 调用商品信息，添加商品，，，选择类目的接口
	//	content.ContentHandler() // 调用内容信息
	s.ListenAndServe()
	fmt.Println("loaclhost:8089")
}
*/
