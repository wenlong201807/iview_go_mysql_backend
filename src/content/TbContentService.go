package content

import (
	"commons"
	"fmt"
	"strconv"
	"strings"
)

// 内容页面分页查询
func showContentService(page, rows int) (e *commons.Datagrid) {
	ts := contentByPageDao(rows, page) // 注意参数顺序，形参与实参位置相对应
	if ts != nil {
		// e 默认是空指针，需要使用new创建 否则报错  nil pointer dereference
		e = new(commons.Datagrid)
		e.Rows = ts          //设置当前页显示数据
		e.Total = contentCountDao() // 数据表中总条数
		return
	}
	return nil
}

// 删除内容--支持同时删多条多个id英文逗号隔开的字符串
func delByIdsService(ids string) (e commons.EgoResult) {
	idStr := strings.Split(ids, ",")
	idInt := make([]int, 0)
	for _, n := range idStr {
		id, _ := strconv.Atoi(n)
		idInt = append(idInt, id)
	}
	count := delByIdsDao(idInt)
	if count > 0 {
		e.Status = 200
		e.Msg = "删除成功"
	}
	return
}

// 新增规格参数**单表添加，最简单
func insertContentService(categoryId int, title,subTitle,titleDesc,url,pic,pic2,content string) (e commons.EgoResult) {
	//date := time.Now().Format("2006-01-02 15:04:05") // 格式化时间方式
	//param := TbContent{ItemCatId: catid, ParamData: paramData, Created: date, Updated: date}
	param := TbContent{CategoryId:categoryId,Title:title,SubTitle:subTitle,TitleDesc:titleDesc,Url:url,Pic:pic,Pic2:pic2,Content:content}
	count := insertContentDao(param)
	if count > 0 {
		e.Status = 200
		e.Msg = "添加成功"
	}
	return
}

// 编辑内容
func updateContentService(id,categoryId int, title,subTitle,titleDesc,url,pic,pic2,content string) (e commons.EgoResult) {
	param := TbContent{Id:id,CategoryId:categoryId,Title:title,SubTitle:subTitle,TitleDesc:titleDesc,Url:url,Pic:pic,Pic2:pic2,Content:content}
	fmt.Println("updateContentService--param:",param)
	count := updContentByIdDao(param)
	fmt.Println("updateContentService---count:",count)
	if count > 0 {
		e.Status = 200
		e.Msg = "修改成功"
	}
	return
}