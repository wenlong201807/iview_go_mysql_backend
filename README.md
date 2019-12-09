#   iview_go_mysql_backEnd项目介绍

-   [本地位置](D:\gitCode\golang_ego)


-   D:\gitCode\golang_ego\src\item\cat   TbItemService.go文件中有一个错误，标注：组合添加失败

-   类似的正确的写法有一个  标注： 修改页面显示信息TbItemService.go
-   D:\gitCode\golang_ego\src\item


##(难点)

-  修改数据接口，设计三个表之间的插入，属于表结构中的事务***重点视频17开始处
-  规格参数，修改接口，删除接口为单表独立操作
-  新增商品接口，新添加另一张表tb_item_param_item中一个字段paramData

-  查询接口出问题，没报错
-  tb_item_param_item接口查询paramData没有正常返回？？？===已解决
-  没有正常启动服务

-   两表联查
-   规格参数页面分页显示--出问题
-   catName 字段属于另一张表的字段，没有正确关联id值，出问题

-  paramData字段，对应页面中，产品规格
-  [{"group":"主体1","params":[{"k":"品牌1","v":"华为1(HUAWEI)"},{"k":"品牌1","v":"华为1(HUAWEI)"}]},{"group":"主体2","params":[{"k":"品牌2","v":"华为1(HUAWEI)"},{"k":"品牌2","v":"华为2(HUAWEI)"}]}]



###  三表关联添加操作成功 ***难点
-   tb_item_param_item
-   tb_item
-   tb_item_desc

-   对应的修改操作失败，使用到事务