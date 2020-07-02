package whysql

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

func GetWhereSqlOrderLimt(TabName, ParameterStr string, OrderStr string, SortStr string, PageSize, CurrentPage int) (string, error) {

	whereSql, err := GetWhereSql(ParameterStr)
	if err != nil {
		return "", err
	}

	whereSql = "select * from  " + TabName + "  where " + whereSql + " Order By " + OrderStr + " " + SortStr + "  LIMIT " + strconv.Itoa((CurrentPage-1)*PageSize) + "," + strconv.Itoa(PageSize)
	return whereSql, nil
}
func GetWhereSqlLimt(TabName, ParameterStr string, PageSize, CurrentPage int) (string, error) {

	whereSql, err := GetWhereSql(ParameterStr)
	if err != nil {
		return "", err
	}

	whereSql = "select * from  " + TabName + "  where " + whereSql + "  LIMIT " + strconv.Itoa((CurrentPage-1)*PageSize) + "," + strconv.Itoa(PageSize)
	return whereSql, nil
}
func GetWhereSqlCount(TabName, ParameterStr string) (string, error) {

	whereSql, err := GetWhereSql(ParameterStr)
	if err != nil {
		return "", err
	}

	whereSql = "select count(1) as Num from  " + TabName + "  where " + whereSql
	return whereSql, nil
}
func GetWhereSql(ParameterStr string) (string, error) {
	sqlWhere := " 1=1 "
	if ParameterStr == "" || ParameterStr == "[]" {
		return sqlWhere, nil
	}
	var filterModelList []FilterModel

	err := json.Unmarshal([]byte(ParameterStr), &filterModelList)
	if err != nil {
		return "", err
	}
	for i := 0; i < len(filterModelList); i++ {
		fieldWhere, err := getFieldWhere(&(filterModelList[i]))
		if err != nil {
			return "", err
		}
		if sqlWhere == "" {
			sqlWhere = "1=1 " + fieldWhere
		} else {
			sqlWhere += fieldWhere
		}

	}
	return sqlWhere, nil
}
func getFieldWhere(model *FilterModel) (string, error) {
	strTemp := ""
	if model == nil {
		return "", errors.New("对象为空")
	}
	if (*model).Logic == "" {
		return "", errors.New("关系符为空")
	}
	if (*model).Action == "" {
		return "", errors.New("算术符为空")
	}
	strTemp = " " + (*model).Logic
	strTemp += " " + (*model).Column + " " + (*model).Action
	strTemp += getwhereByDataType((*model).DataType)
	strTemp += getwhereByAction((*model).Action)
	strTemp += (*model).Value
	strTemp += getwhereByAction((*model).Action)
	strTemp += getwhereByDataType((*model).DataType)

	return strTemp, nil

}
func getwhereByAction(action string) string {
	action = strings.ToLower(action)
	switch action {
	case "like":
		return "%"
	default:
		return ""
	}
}

func getwhereByDataType(dataType string) string {
	switch dataType {
	case "S", "D":
		return "'"
	case "I", "F":
		return ""
	default:
		return "'"
	}
}

func DelSqlByField(TabName, FieldName string, Value interface{}) (string, error) {
	var err error
	var Sql string
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("类型断言错误")
		}
	}()
	if Value == nil {
		return "", errors.New("值为空")
	}
	Sql = " delete FROM " + TabName + " where  " + FieldName + " = "
	Sql1 := ""
	switch Value.(type) {
	case int:
		Sql1 = strconv.Itoa(Value.(int))
	case int64:
		Sql1 = strconv.FormatInt(Value.(int64), 10)
	case float64:
		Sql1 = strconv.FormatFloat(Value.(float64), 'E', -1, 64)
	case string:
		Sql1 = "'" + Value.(string) + "'"
	default:
		Sql1 = ""
	}
	Sql = Sql + Sql1
	return Sql, err
}
