package cat

import (
	"commons"
	"fmt"
)

//根据id查询类目
func ShowCatByIdService(id int) *TbItemCat  {
	fmt.Println("ShowCatByIdService--id:",id)
	return selByIdDao(id)
}

// 查询商品中添加商品的----选择类目部分----树状结构处理
func showCatByPidService(pid int) (tree []commons.EasyUITree)  {
	cats := selByPidDao(pid)
	tree = make([]commons.EasyUITree,0) // 必须先实例化一个空的
	for _,n := range cats{
		state := "open" // 不能使用单引号 // tree 树状结构页面展示默认打开(展开状态)
		if n.IsParent{
			state = "close"
		}
		tree = append(tree,commons.EasyUITree{n.Id,n.Name,state})
	}
	return

}