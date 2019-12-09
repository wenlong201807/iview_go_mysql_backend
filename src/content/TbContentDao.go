package content

import (
	"commons"
	"fmt"
	"strconv"
)

/*
rows:每页显示条数
page:当前第几页
*/
// 查询数据
func contentByPageDao(rows, page int) []TbContent {
	// 第一个表示：从哪条开始查询，0算起，第二个：查询几个
	r, err := commons.Dql("select * from tb_content limit ?,?", rows*(page-1), rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	ts := make([]TbContent, 0)
	for r.Next() {
		var t TbContent
		// 当某个值为null时，数据会录入失败，需要注意 Barcode
		//var s sql.NullString
		//r.Scan(&t.Id, &t.Title, &t.Content, &t.Created, &t.Updated)
		r.Scan(&t.Id, &t.CategoryId,&t.Title, &t.SubTitle, &t.TitleDesc, &t.Url, &t.Pic, &t.Pic2, &t.Content, &t.Created, &t.Updated)

		//t.Barcode = s.String
		ts = append(ts, t)
	}
	commons.CloseConn()
	return ts
}

// 查询总条数***正确  // 如果返回值为<0，表示查询失败
func contentCountDao() (count int) {
	rows, err := commons.Dql("select count(*) from tb_content")
	if err != nil {
		fmt.Println(err)
		return -1 // 返回值为<0，表示查询失败
	}
	rows.Next() // 查询成功时，数据必然是一行一列
	rows.Scan(&count)
	commons.CloseConn() // 关闭连接
	return
}

// 支持多条同时删除
func delByIdsDao(ids []int) int {
	sql := "delete from tb_content where id in ("
	for i:=0;i<len(ids);i++{
		sql += strconv.Itoa(ids[i]) // int--> string
		if i < len(ids) -1 { // 最后一个不要逗号
			sql += ","
		}
	}
	sql += ")"
	count,err := commons.Dml(sql)
	if err != nil{
		fmt.Println(err)
		return -1
	}
	return int(count)
}

// 新增
func insertContentDao(param TbContent) int {
	sql := "insert into tb_content values(default,?,?,?,?,?,?,?,?,now(),now())"
	count,err := commons.Dml(sql,param.CategoryId,param.Title,param.SubTitle,param.TitleDesc,param.Url,param.Pic,param.Pic2,param.Content)
	if err != nil{
		fmt.Println(err)
		return -1
	}
	return int(count)
}

//修改
func updContentByIdDao(param TbContent) int {
	sql := "update tb_content set category_id=?,title=?,sub_title=?,title_desc=?,url=?,pic=?,pic2=?,content=?,updated=now() where id=?"
	count,err := commons.Dml(sql,param.CategoryId,param.Title,param.SubTitle,param.TitleDesc,param.Url,param.Pic,param.Pic2,param.Content,param.Id)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	fmt.Println("updContentByIdDao--count:",count)
	return int(count)
}