package cat

import (
	"commons"
	"fmt"
)

func selByIdDao(id int) (t *TbItemCat) {
	rows, err := commons.Dql("select * from tb_item_cat where id=?", id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if rows.Next() {// 如果有数据，就进行实例化填充，再进行关闭和返回
		t = new(TbItemCat)
		rows.Scan(&t.Id, &t.ParentId, &t.Name, &t.Status, &t.SortOrder, &t.IsParent, &t.Created, &t.Updated)
	}
	// 如果没有，就只要关闭，和返回
	commons.CloseConn()
	fmt.Println("TbItemCatDao--selByIdDao--id:",id)
	return t
}

/*
根据parent_id查询所有子类目
*/
func selByPidDao(pid int) (c []TbItemCat) { // []TbItemCat 有[] 表示多行数据 * 表示指针类型
	rows,err := commons.Dql("select * from tb_item_cat where parent_id=?",pid)
	if err != nil{
		fmt.Println(err)
		return
	}
	c = make([]TbItemCat,0) // 传入的是指针类型，必须要实例化一个，否则报错
	for rows.Next(){
		var t TbItemCat
		rows.Scan(&t.Id,&t.ParentId,&t.Name,&t.Status,&t.SortOrder,&t.IsParent,&t.Updated,&t.Created)
		c = append(c,t)
	}
	commons.CloseConn() // 关闭连接
	return
}
