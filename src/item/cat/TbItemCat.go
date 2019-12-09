package cat

// 商品类目
type TbItemCat struct {
	Id        int
	ParentId  int
	Name      string
	Status    string
	SortOrder int8
	IsParent  bool  // byte // 数据表中定义的数据类型为tinyint,与服务层(TbItemCatService.go)不相符
	Created   string
	Updated   string
}
