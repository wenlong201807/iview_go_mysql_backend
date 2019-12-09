package commons

import "github.com/gorilla/mux"

var (
	Router = mux.NewRouter()
	CurrPath = "http://localhost:80/"
	HEADER_CONTENT_TYPE = "Content-Type"
	JSON_HEADER = "application/json;charset=utf-8"
)
