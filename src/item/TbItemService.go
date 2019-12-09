package item

import (
	"commons"
	"fmt"
	"io/ioutil"
	"item/cat"
	"item/desc"
	"item/paramitem"
	"math/rand"
	"mime/multipart"
	"net/url"
	"strconv"
	"strings"
	"time"
)

/*
// 父类目的name组合添加失败
func showItemService(page,rows int) (e *commons.Datagrid)  {
	ts := selByPageDao(rows,page) // 注意参数顺序，
	if ts != nil{
		// 新添加部分
		itemChildren := make([]TbItemChild,0)
		for i:=0;i<len(ts);i++{
			var itemChild TbItemChild
			itemChild.Id = ts[i].Id
			itemChild.Updated = ts[i].Updated
			itemChild.Created = ts[i].Created
			itemChild.Status = ts[i].Status
			itemChild.Barcode = ts[i].Barcode
			//itemChild.Cid = ts[i].Cid // 未使用
			//itemChild.Image = ts[i].Image // 未使用
			itemChild.Num = ts[i].Num
			itemChild.Price = ts[i].Price
			itemChild.SellPoint = ts[i].SellPoint
			itemChild.Title = ts[i].Title
			//fmt.Print("===",cat.ShowCatByService(ts[i].Cid).Name)
			// 额外添加的一个
			itemChild.CategoryName = cat.ShowCatByIdService(ts[i].Cid).Name // 报错行***函数名调用出错，
			//数据转换完成丢到一个数组中
			itemChildren = append(itemChildren,itemChild)
		}



		// e 默认是空指针，需要使用new创建 否则报错  nil pointer dereference
		e = new(commons.Datagrid)
		// e.Rows = ts //设置当前页显示数据 // 原有数据返回给前端
		e.Rows = itemChildren // 重新组合的数据返回给前端
		e.Total = selCount() // 数据表中总条数
		return
	}
	return nil
}
*/

// 没有子类目的正常代码
func showItemService(page, rows int) (e *commons.Datagrid) {
	ts := selByPageDao(rows, page) // 注意参数顺序，
	if ts != nil {
		// e 默认是空指针，需要使用new创建 否则报错  nil pointer dereference
		e = new(commons.Datagrid)
		e.Rows = ts          //设置当前页显示数据
		e.Total = selCount() // 数据表中总条数
		return
	}
	return nil
}

// 删除商品
func delByIdsService(ids string) (e commons.EgoResult) {
	count := updStatusByIdsDao(strings.Split(ids, ","), 3) // 删除标志为字符串3
	fmt.Print("删除商品条数count:", count)
	if count > 0 {
		e.Status = 200
		e.Msg = "商品删除成功，对应的状态值修改为3"

	}
	return
}

// 商品上架
func instockService(ids string) (e commons.EgoResult) {
	count := updStatusByIdsDao(strings.Split(ids, ","), 1)
	if count > 0 {
		e.Status = 200
		e.Msg = "商品上架成功，对应的状态值修改为1"
	}
	return
}

// 商品下架
func offStockService(ids string) (e commons.EgoResult) {
	count := updStatusByIdsDao(strings.Split(ids, ","), 2)
	if count > 0 {
		e.Status = 200
		e.Msg = "商品下架成功，对应的状态值修改为2"
	}
	return
}

// 上传图片
func imageUplaodService(f multipart.File, h *multipart.FileHeader) map[string]interface{} { // 返回值有两种可能，不能是结构体，只能使用数组，自定义
	m := make(map[string]interface{}) // 实例化一个
	b, err := ioutil.ReadAll(f)
	if err != nil {
		m["error"] = 1
		m["message"] = "上传失败，服务器错误"
		return m
	}
	// 成功时，写入图片路径 组成部分：纳秒时间戳+随机数+扩展名
	rand.Seed(time.Now().UnixNano()) // 设置种子？为什么要这个
	fileName := "static/images" + strconv.Itoa(int(time.Now().UnixNano())) + strconv.Itoa(rand.Intn(1000)) + h.Filename[strings.LastIndex(h.Filename, "."):]
	err = ioutil.WriteFile(fileName, b, 0777)
	if err != nil {
		m["error"] = 1
		m["message"] = "上传失败，保存图片时错误"
		return m
	}
	m["error"] = 0
	m["url"] = commons.CurrPath + fileName
	return m
}

// 商品新增 ** 同一个前端数据传入后端的到两张表中
func insertService(f url.Values) (e commons.EgoResult) {
	var t TbItem
	cid, _ := strconv.Atoi(f["Cid"][0])
	t.Cid = cid
	t.Title = f["Title"][0]
	t.SellPoint = f["SellPoint"][0]
	price, _ := strconv.Atoi(f["Price"][0])
	t.Price = price
	num, _ := strconv.Atoi(f["Num"][0])
	t.Num = num
	t.Image = f["Image"][0]
	t.Status = 1
	date := time.Now().Format("2006-01-02 15:04:05")
	t.Created = date
	t.Updated = date
	id := commons.GenId()
	t.Id = id
fmt.Println("新添商品--insertService--t TbItem:",t)
	//t TbItem: {19031574842314 标题党99 卖点99 99900 99  99 99 1 2019-11-27 16:11:54 2019-11-27 16:11:54}
	// 商品表新增
	count := insertItemDao(t)
	if count > 0 { //新增成功的话，需要插入其他表中的字段

		var tbItemDesc desc.TbItemDesc
		tbItemDesc.ItemId = id
		tbItemDesc.Created = date
		tbItemDesc.Updated = date
		tbItemDesc.ItemDesc = f["Desc"][0]
		fmt.Println("新添商品--insertService--tbItemDesc desc.TbItemDesc:",tbItemDesc)
		//tbItemDesc desc.TbItemDesc: {19031574842314 主备999用来修改的，先在这添加 2019-11-27 16:11:54 2019-11-27 16:11:54}
		countDesc := desc.Insert(tbItemDesc)
		if countDesc > 0 { // 新增商品成功，必须同时保证数据可以传入到两种表中，tb_item,tb_item_desc
			//新添加规格参数表中的商品规格信息
			paramItem := paramitem.TbItemParamItem{ItemId: id, ParamData: f["paramData"][0]}
			fmt.Println("新添商品--insertService-- paramitem TbItemParamItem:",paramItem)
			//paramitem TbItemParamItem: {0 19031574842314 [{"group":"主体99","params":[{"k":"品牌1","v":"华为1(HUAWEI)"},
			// {"k":"品牌1","v":"华为1(HUAWEI)"}]},
			// {"group":"主体2","params":[{"k":"品牌2","v":"华为1(HUAWEI)"},{"k":"品牌2","v":"华为2(HUAWEI)"}]}]  }
			countParamItem := paramitem.InsertParamItemDao(paramItem)
			if countParamItem > 0 {
				e.Status = 200
				e.Msg ="添加成功"
			} else { // 删除商品中的数据，同时删除商品描述中的数据
				// tb_item表中数据插入成功，但是tb_item_desc表中插入失败
				// 因此，需要将tb_item插入插入成功的数据删掉，保证表中的数据统一
				delById(id)
				desc.DelDescByIdDao(id)
				e.Status = 400
			}

		} else {
			// tb_item表中数据插入成功，但是tb_item_desc表中插入失败
			// 因此，需要将tb_item插入插入成功的数据删掉，保证表中的数据统一
			delById(id)
			e.Status = 400
		}
	}
	return
}

// 修改页面显示信息
func showItemDescCatService(id int) TbItemDescChild {
	item := selByIdDao(id)
	var c TbItemDescChild
	c.Id = item.Id
	c.Updated = item.Updated
	c.Created = item.Created
	c.Barcode = item.Barcode
	c.Status = item.Status
	c.Price = item.Price
	c.Num = item.Num
	c.Title = item.Title
	c.Cid = item.Cid
	c.Image = item.Image
	c.SellPoint = item.SellPoint
	// 商品类目
	c.CategoryName = cat.ShowCatByIdService(c.Cid).Name
	// 商品描述
	c.Desc = desc.SelByIdService(c.Id).ItemDesc
	return c
}

// 修改商品
func updateService(v url.Values) (e commons.EgoResult) {
	commons.OpenConnWithTx()
	var t TbItem
	id, _ := strconv.Atoi(v["Id"][0])
	fmt.Println("TbItemService--updateService--id:",id)
	t.Id = id
	cid, _ := strconv.Atoi(v["Cid"][0])
	t.Cid = cid
	t.Title = v["Title"][0]
	t.SellPoint = v["SellPoint"][0]
	price, _ := strconv.Atoi(v["Price"][0])
	t.Price = price
	num, _ := strconv.Atoi(v["Num"][0])
	t.Num = num
	t.Image = v["Image"][0]
	status, _ := strconv.Atoi(v["Status"][0])
	t.Status = int8(status)
	date := time.Now().Format("2006-01-02 15:04:05")
	t.Updated = date
	count := updateItemByIdWithTxDao(t)

	if count > 0 {
		var itemDesc desc.TbItemDesc
		itemDesc.ItemId = id
		itemDesc.ItemDesc = v["Desc"][0]
		itemDesc.Updated = date
		count = desc.UpdateDescByIdWithTxDao(itemDesc)
		if count > 0 {
			// 修改了商品和描述后，再修改商品的规格参数**新加的
			count = paramitem.UpdByItemIdWithTxDao(id,v["paramData"][0])
			if count > 0 {
				commons.CloseConnWithTx(true)
				e.Status = 200
				return
			}

		}
	}
	commons.CloseConnWithTx(false)
	return
}
