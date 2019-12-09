package commons

// 对应  查询商品功能  表  tb_item  添加  表 tb_content
type Datagrid struct {
	// 当前页显示的数据
	Rows interface{} `json:"rows"`
	// 总个数
	Total int `json:"total"`
}

// 依据easyui  tree的数据结构特点，构造数据结构
type EasyUITree struct {
	Id int `json:"id"`
	Text string `json:"text"`
	State string `json:"state"`
} 
