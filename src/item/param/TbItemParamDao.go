package param

import (
	"commons"
	"fmt"
	"strconv"
)

// 分页查询
func selByPageDao(page, rows int) []TbItemParam {
	r, err := commons.Dql("select * from tb_item_param limit ?,?", rows*(page-1), rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	t := make([]TbItemParam, 0)
	for r.Next() {
		var param TbItemParam
		r.Scan(&param.Id, &param.ItemCatId, &param.ParamData, &param.Created, &param.Updated)
		t = append(t, param)
	}
	return t
}

//查询总条数
func selCountDao() int {
	r, err := commons.Dql("select count(*) from tb_item_param")
	if err != nil {
		fmt.Println(err)
		return -1
	}
	if r.Next() {
		var count int
		r.Scan(&count)
		return count
	}
	return -1
}

// 删除多条数据
func delByIdsDao(ids []int) int {
	sql := "delete from tb_item_param where id in ("
	for i := 0; i < len(ids); i++ {
		sql += strconv.Itoa(ids[i]) // int --> string
		if i < len(ids)-1 {
			sql += ","
		}
	}
	sql += ")"
	count, err := commons.Dml(sql)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return int(count)
}

// 根据类目id查询规格参数
func selByCatIdDao(catid int) *TbItemParam {
	r, err := commons.Dql("select * from tb_item_param where item_cat_id=?", catid)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if r.Next() {
		param := new(TbItemParam)
		r.Scan(&param.Id, &param.ItemCatId, &param.ParamData, &param.Created, &param.Updated)
		return param
	}
	return nil
}

// 新增
func insertParamDao(param TbItemParam) int {
	count,err := commons.Dml("insert into tb_item_param values(default,?,?,?,?)",param.ItemCatId,param.ParamData,param.Created,param.Updated)
	if err != nil{
		fmt.Println(err)
		return -1
	}
	return int(count)
}

//修改
func updParamByIdDao(param TbItemParam) int {
	count,err := commons.Dml("update tb_item_param set item_cat_id=?,param_data=?,updated=now() where id=?",param.ItemCatId,param.ParamData,param.Id)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return int(count)
}
