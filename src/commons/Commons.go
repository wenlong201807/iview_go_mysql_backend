package commons

import (
	"math/rand"
	"strconv"
	"time"
)

// 生成数据库的主键：自动生成的id值，时间纳米数与毫秒数拼接而成
func GenId() int {
	rand.Seed(time.Now().UnixNano())
	id, _ := strconv.Atoi(strconv.Itoa(rand.Intn(10000)) + strconv.Itoa(int(time.Now().Unix())))
	return id
}
