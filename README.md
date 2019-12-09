# iview_go_mysql_backend 项目介绍

- [本地位置后端](D:\gitCode\golang_ego)
- [本地位置前端](D:\gitCode\vue\iview_go_mysql)

- [厂库前端地址](https://github.com/wenlong201807/iview_go_mysql_frontend)

* D:\gitCode\golang_ego\src\item\cat TbItemService.go 文件中有一个错误，标注：组合添加失败

* 类似的正确的写法有一个 标注： 修改页面显示信息 TbItemService.go
* D:\gitCode\golang_ego\src\item

##(难点)

- 修改数据接口，设计三个表之间的插入，属于表结构中的事务\*\*\*重点视频 17 开始处
- 规格参数，修改接口，删除接口为单表独立操作
- 新增商品接口，新添加另一张表 tb_item_param_item 中一个字段 paramData

- 查询接口出问题，没报错
- tb_item_param_item 接口查询 paramData 没有正常返回？？？===已解决
- 没有正常启动服务

- 两表联查
- 规格参数页面分页显示--出问题
- catName 字段属于另一张表的字段，没有正确关联 id 值，出问题

- paramData 字段，对应页面中，产品规格
- [{"group":"主体 1","params":[{"k":"品牌 1","v":"华为 1(HUAWEI)"},{"k":"品牌 1","v":"华为 1(HUAWEI)"}]},{"group":"主体 2","params":[{"k":"品牌 2","v":"华为 1(HUAWEI)"},{"k":"品牌 2","v":"华为 2(HUAWEI)"}]}]

### 三表关联添加操作成功 \*\*\*难点

- tb_item_param_item
- tb_item
- tb_item_desc

- 对应的修改操作失败，使用到事务
