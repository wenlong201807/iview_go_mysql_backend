package user

import (
	"commons"
	"encoding/json"
	"fmt"
	"net/http"
)

// 所有user模块的handler
func UserHandler() {
	commons.Router.HandleFunc("/login", loginController)
	//http.HandleFunc("/login",loginController) // 接口为login的请求
}

// 登录
func loginController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//fmt.Println("PATH: ", r.URL.Path)
	//fmt.Println("SCHEME: ", r.URL.Scheme)
	//fmt.Println("METHOD: ", r.Method)
	//fmt.Println("Host: ", r.Host)
	//fmt.Println("RequestURI: ", r.RequestURI)
	//fmt.Println("RemoteAddr: ", r.RemoteAddr)
	//fmt.Println("Form: ", r.Form)
	//fmt.Println("PostForm: ", r.PostForm)
	//fmt.Println("PostForm['username']: ", r.PostForm["username"])
	/*
		PATH:  /login
		SCHEME:
		METHOD:  POST
		Host:  192.168.43.148:80
		RequestURI:  /login?username=bjsxt&password=smallming
		RemoteAddr:  192.168.43.148:51814
		Form:  map[password:[smallming] username:[bjsxt]]
		PostForm:  map[]
	*/
	//username := r.FormValue("username")
	//password := r.FormValue("password")
	// post请求接受参数方式可行一
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	fmt.Println("loginController--usrname,password==post:",username,password)
	er := LoginService(username, password)
	// 把结构体转换为json数据
	b, _ := json.Marshal(er)
	// 设置响应内容为json
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)

}
