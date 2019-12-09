package commons

// 客户端与服务器段数据交互模板
type EgoResult struct {
	Status int  // 状态，为200表示成功，其他为失败
	Data interface{} // 返回的数据
	Msg string // 返回的消息
}