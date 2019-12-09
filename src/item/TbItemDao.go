package item

import (
	"commons"
	"database/sql"
	"fmt"
)

/*
rows:每页显示条数
page:当前第几页
*/

func selByPageDao(rows, page int) []TbItem {
	// 第一个表示：从哪条开始查询，0算起，第二个：查询几个
	r, err := commons.Dql("select * from tb_item limit ?,?", rows*(page-1), rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	ts := make([]TbItem, 0)
	for r.Next() {
		var t TbItem
		// 当某个值为null时，数据会录入失败，需要注意 Barcode
		var s sql.NullString
		r.Scan(&t.Id, &t.Title, &t.SellPoint, &t.Price, &t.Num, &s, &t.Image, &t.Cid, &t.Status, &t.Created, &t.Updated)
		//r.Scan(&t.Id,&t.Title,&t.SellPoint,&t.Price,&t.Num,&t.Barcode,&t.Image,&t.Cid,&t.Status,&t.Created,&t.Updated)
		t.Barcode = s.String
		ts = append(ts, t)
	}
	commons.CloseConn()
	return ts

}

// 查询总条数
// 如果返回值为<0，表示查询失败
func selCount() (count int) {
	rows, err := commons.Dql("select count(*) from tb_item")
	if err != nil {
		fmt.Println(err)
		return -1 // 返回值为<0，表示查询失败
	}
	rows.Next() // 查询成功时，数据必然是一行一列
	rows.Scan(&count)
	commons.CloseConn() // 关闭连接
	return
}

/*
返回值如果小于0，表示更新失败
*/
func updStatusByIdsDao(ids []string, status int) int {
	if len(ids) <= 0 {
		return -1
	}
	//sql := "update tb_item set status=? where id=? or id=?" // 拼接数据格式，最后一个id不要or连接
	sql := "update tb_item set status=? where "
	for i := 0; i < len(ids); i++ {
		sql += " id=" + ids[i]
		if i < len(ids)-1 {
			sql += " or "
		}
	}
	count, err := commons.Dml(sql, status)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return int(count)
}

// 商品新增
func insertItemDao(t TbItem) int {
	count, err := commons.Dml("insert into tb_item values(?,?,?,?,?,?,?,?,?,?,?)", t.Id, t.Title, t.SellPoint, t.Price, t.Num, t.Barcode, t.Image, t.Cid, t.Status, t.Created, t.Updated)
	if err != nil {
		return -1
	}
	return int(count)
}

// 根据id删除
func delById(id int) int {
	count, err := commons.Dml("delete from tb_item where id=?", id)
	if err != nil {
		fmt.Println("根据id删除的错误信息：", err)
		return -1
	}
	return int(count)
}

// 修改前的查询，依据id查询后端表中的数据，返回给修改前的修改模态框的数据
// 根据主键查询内容
func selByIdDao(id int) *TbItem { // dml，dql之间的差别？？
	rows, err := commons.Dql("select * from tb_item where id=?", id)
	if err != nil {
		fmt.Println("修改前的查询err:", err)
		return nil
	}
	if rows.Next() {
		t := new(TbItem)
		var s sql.NullString
		rows.Scan(&t.Id, &t.Title, &t.SellPoint, &t.Price, &t.Num, &s, &t.Image, &t.Cid, &t.Status, &t.Created, &t.Updated)
		t.Barcode = s.String
		return t
	}
	return nil
}

// 修改商品表数据
func updateItemByIdWithTxDao(t TbItem) int  {
	return commons.PrepareWithTx("update  tb_item set title=?,sell_point=?,price=?,num=?,barcode=?,image=?,cid=?,status=?,updated=? where id=?",
		t.Title,t.SellPoint,t.Price,t.Num,t.Barcode,t.Image,t.Cid,t.Status,t.Updated,t.Id)
}