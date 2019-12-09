package item

import (
	"commons"
	"encoding/json"
	"net/http"
	"strconv"
)

// 接口调用函数// 查询商品
func ItemHandler() {
	commons.Router.HandleFunc("/showItem", showItemController)
	commons.Router.HandleFunc("/item/delete", delByIdsController)
	commons.Router.HandleFunc("/item/instock", instockController)
	commons.Router.HandleFunc("/item/offstock", offstockController)
	commons.Router.HandleFunc("/item/imageupload", imagesUploadController)
	commons.Router.HandleFunc("/item/add", insertController)
	commons.Router.HandleFunc("/item/showItemById", showItemDescCatController) //显示修改前的信息
	commons.Router.HandleFunc("/item/update", updateController)
}

// 修改商品
func updateController(w http.ResponseWriter ,r *http.Request)  {
	r.ParseForm()
	er := updateService(r.Form)
	b, _ := json.Marshal(er)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)

}

// 修改前，显示的要修改的信息
func showItemDescCatController(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	c := showItemDescCatService(id)
	b, _ := json.Marshal(c)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}

// 商品新增
func insertController(w http.ResponseWriter, r *http.Request) {
	// 获取的参数需要先解析，否则为空值
	r.ParseForm()
	er := insertService(r.Form)
	b, _ := json.Marshal(er)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}

// 图片上传 *** 没有实测
func imagesUploadController(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("imgFile")
	if err != nil {
		m := make(map[string]interface{})
		m["error"] = 1
		m["message"] = "接收图片失败"
		b, _ := json.Marshal(m)
		//w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
		w.Write(b)
		return
	}
	// 图片上传成功时的效果
	m := imageUplaodService(file, fileHeader)
	b, _ := json.Marshal(m)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}

// 显示商品信息
func showItemController(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	rows, _ := strconv.Atoi(r.FormValue("rows"))
	datagrid := showItemService(page, rows)
	// 转换为json
	b, _ := json.Marshal(datagrid)
	// json格式必须设置响应头的东西
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}

// 商品删除
func delByIdsController(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")
	er := delByIdsService(ids)
	b, _ := json.Marshal(er)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}

// 商品上架
func instockController(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")
	er := instockService(ids)
	b, _ := json.Marshal(er)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}

// 商品下架
func offstockController(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")
	er := offStockService(ids)
	b, _ := json.Marshal(er)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}
