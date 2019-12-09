package desc

func Insert(t TbItemDesc) int  {
	return insertDescDao(t)
}

// 依据id查询tb_item_desc
func SelByIdService(id int) *TbItemDesc  {
	return selByIdDao(id)
}