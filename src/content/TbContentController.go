package content

import (
	"commons"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)


// 接口调用函数// 查询内容
func ContentHandler() {
	commons.Router.HandleFunc("/showContent", showContentController)
	commons.Router.HandleFunc("/deleteContent", delByIdsController)
	commons.Router.HandleFunc("/addContent", insertContentController)
	commons.Router.HandleFunc("/updateContent", updateContentController)

}


// 显示分页内容信息
func showContentController(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	rows, _ := strconv.Atoi(r.FormValue("rows"))
	datagrid := showContentService(page, rows)
	// 转换为json
	b, _ := json.Marshal(datagrid)
	//fmt.Println("showContentController返回前端的内容",b)
	// json格式必须设置响应头的东西
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}

//删除
func delByIdsController(w http.ResponseWriter, r *http.Request) {
	er := delByIdsService(r.FormValue("ids")) // 前端使用的参数
	fmt.Println("delByIdsController--delete id:",r.FormValue("ids"))
	b, _ := json.Marshal(er)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}

// 新增
func insertContentController(w http.ResponseWriter, r *http.Request)  {
	categoryId,_ := strconv.Atoi(r.PostFormValue("categoryId"))
	title := r.PostFormValue("title")
	fmt.Println("新增，insertContentController-PostFormValue--title:",title)
	subTitle := r.PostFormValue("subTitle")
	titleDesc := r.PostFormValue("titleDesc")
	url := r.PostFormValue("url")
	pic := r.PostFormValue("pic")
	pic2 := r.PostFormValue("pic2")
	content := r.PostFormValue("content")
	er := insertContentService(categoryId, title,subTitle,titleDesc,url,pic,pic2,content)
	b, _ := json.Marshal(er)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}

// 修改
func updateContentController(w http.ResponseWriter, r *http.Request)  {


	categoryId,_ := strconv.Atoi(r.FormValue("categoryId"))
	Id,_ := strconv.Atoi(r.FormValue("id")) // 不可修改的，索引id
	fmt.Println("修改，categoryId:",categoryId)
	fmt.Println("修改，索引id:",Id)
	title := r.FormValue("title")
	subTitle := r.FormValue("subTitle")
	titleDesc := r.FormValue("titleDesc")
	url := r.FormValue("url")
	pic := r.FormValue("pic")
	pic2 := r.FormValue("pic2")
	content := r.FormValue("content")
	er := updateContentService(Id,categoryId, title,subTitle,titleDesc,url,pic,pic2,content)
	b, _ := json.Marshal(er)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}