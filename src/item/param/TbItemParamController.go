package param

import (
	"commons"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func ParamHandler() {
	commons.Router.HandleFunc("/item/param/show", showParamController)
	commons.Router.HandleFunc("/item/param/delete", delByIdsController)
	commons.Router.HandleFunc("/item/param/iscat", isCatController)
	commons.Router.HandleFunc("/item/param/add", insertParamController)
	commons.Router.HandleFunc("/item/param/edit", updateParamController)
}

// 编辑规格参数
func updateParamController(w http.ResponseWriter, r *http.Request)  {
	id,_ := strconv.Atoi(r.FormValue("id"))
	itemCatId,_ := strconv.Atoi(r.FormValue("itemCatId"))
	paramData := r.FormValue("paramData")
	er := updateParamService(id,itemCatId,paramData)
	b, _ := json.Marshal(er)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}

// 新增
func insertParamController(w http.ResponseWriter, r *http.Request)  {
	catid,_ := strconv.Atoi(r.FormValue("itemCatId"))
	paramData := r.FormValue("paramData")
	er := insertParamService(catid,paramData)
	b, _ := json.Marshal(er)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}

func isCatController(w http.ResponseWriter, r *http.Request)  {
	catid,_ := strconv.Atoi(r.FormValue("catid")) // 变量名字要与表(结构体)中字段对应
	er := catidService(catid)
	b, _ := json.Marshal(er)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}

//删除
func delByIdsController(w http.ResponseWriter, r *http.Request) {
	er := delByIdsService(r.FormValue("ids"))
	b, _ := json.Marshal(er)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}

// 显示规格参数
func showParamController(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	rows, _ := strconv.Atoi(r.FormValue("rows"))
	datagrid := showParamService(page, rows)
	// 转换为json
	b, _ := json.Marshal(datagrid)
	fmt.Println("showContentController返回前端的内容", b)
	// json格式必须设置响应头的东西
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}
