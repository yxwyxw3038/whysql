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
	temp.Column = "Name"
	temp.DataType = "S"
	temp.Logic = "and"
	temp.Value = "yxw"
	filterModelList = append(filterModelList, temp)
	b, err := json.Marshal(filterModelList)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	whereStr := string(b)
	fmt.Print(whereStr)
	whereSql, err := whysql.GetWhereSqlOrderLimt("User", whereStr, "UpdateTime", whysql.DESC, 10, 1)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Println(whereSql)

}
```
输出
```

```json
[{"column":"Name","action":"like","logic":"and","value":"yxw","dataType":"S"}]
```

```sql
select * from  User  where  1=1  and Name like'%yxw%' Order By UpdateTime DESC  LIMIT 0,10
```

## api说明

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

