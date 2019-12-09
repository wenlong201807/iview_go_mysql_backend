package cat

import (
	"commons"
	"encoding/json"
	"net/http"
	"strconv"
)

func ItemCatHandler()  {
	commons.Router.HandleFunc("/item/cat/show", showItemCatController)
}

func showItemCatController(w http.ResponseWriter, r *http.Request)  {
	id := r.FormValue("id")
	if id == ""{
		id = "0"
	}
	// 字符串转int
	idInt,_ := strconv.Atoi(id)
	t := showCatByPidService(idInt)
	b, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}
