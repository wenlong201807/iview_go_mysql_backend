package user

import (
	"commons"
	"fmt"
)

func LoginService(un,pwd string)(er commons.EgoResult)  {
	u := SelByUnPwdDao(un,pwd)
	fmt.Println("LoginService--u:",u)
	if u !=nil{
		er.Status = 200
		er.Msg = "登录成功"

	}else {
		er.Status = 400
	}
	return
}
