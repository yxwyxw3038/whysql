## whysql
mysql SQL语句自动生成器

## 用例
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/yxwyxw3038/whysql"
)

func main() {
	var filterModelList []whysql.FilterModel
	var temp whysql.FilterModel
	temp.Action = "like"
	temp.Column = "AccountName"
	temp.DataType = "S"
	temp.Logic = "and"
	temp.Value = "yxw"
	filterModelList = append(filterModelList, temp)
	b, err := json.Marshal(filterModelList)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	whereStr := string(b)
	fmt.Println(whereStr)
	fmt.Println("函数式")
	whereSql, err := whysql.GetWhereSqlOrderLimt("User", whereStr, "UpdateTime", whysql.DESC, 10, 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(whereSql)
	sqldb, err := whysql.NewWhy(whereStr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("链式")
	whereSql, err = sqldb.SetTabName("User").SetOrderByCustomize(whysql.OrderByModel{Column: "ID", SortType: whysql.DESC}).SetLimt(0, 20).GetQuerySql()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(whereSql)
	whereSql, err = sqldb.SetTabName("User").SetOrderBy("ID", "AccountName", "DESC").SetLimt(0, 20).GetQuerySql()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(whereSql)

}
```
输出
```

```json
[{"column":"AccountName","action":"like","logic":"and","value":"yxw","dataType":"S"}]
```

```sql
select * from  User  where  1=1  and AccountName like'%yxw%' Order By UpdateTime DESC  LIMIT 0,10
```

```sql
select * from User where 1=1  and AccountName like'%yxw%' OrderBy  ID DESC  LIMIT  0,20 
```

```sql
select * from User where 1=1  and AccountName like'%yxw%' OrderBy  ID DESC ,  AccountName DESC  LIMIT  0,20
```
## api说明

```go
func NewWhy(ParameterStr string) (*WhyInfo, error)
```
根据前端的JSON初始化对象


```go
func (m *WhyInfo) Reset() *WhyInfo
```
根据刷新对象


```go
func (m *WhyInfo) Clear() *WhyInfo
```
根据清除对象


```go
func (m *WhyInfo) SetTabName(tabName string) *WhyInfo
```
设表名

```go
func (m *WhyInfo) SetPageSize(pageSize int) *WhyInfo
```
设页面大小

```go
func (m *WhyInfo) SetCurrentPage(currentPage int) *WhyInfo 
```
设当前页

```go
func (m *WhyInfo) SetLimt(args ...int) *WhyInfo
```
设置limt 会重置当前页，页面大小
只传一个值为当前页，页面大小补默认值
只传俩个值为当前页，页面大小

```go
func (m *WhyInfo) SetOrderBy(rlist ...string) *WhyInfo 
```
排序
排序字段，排序方式
只有排序字段,则默认倒序

```go
func (m *WhyInfo) SetOrderByCustomize(rlist ...OrderByModel) *WhyInfo 
```
自定义排序


```go
func (m *WhyInfo) GetQuerySql() (string, error)
```
生成查询语句


```go
func (m *WhyInfo) GetCountSql() (string, error)
```
生成聚合数量查询语句

```go
func GetWhereSqlOrderLimt(TabName, ParameterStr string, OrderStr string, SortStr string, PageSize, CurrentPage int) (string, error)
```
见文思义
```go
func GetWhereSqlLimt(TabName, ParameterStr string, PageSize, CurrentPage int) (string, error)  
```
见文思义  
```go
func GetWhereSqlCount(TabName, ParameterStr string) (string, error) 
```
见文思义
```go
func GetWhereSql(ParameterStr string) (string, error) 
```
见文思义


```go
const (
	ASC  = "ASC"
	DESC = "DESC"
)
```

