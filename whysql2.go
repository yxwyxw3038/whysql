package whysql

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

func NewWhy(ParameterStr string) (*WhyInfo, error) {
	var err error

	defer func() {
		if p := recover(); p != nil {
			err = errors.New("初始化异常")
		}
	}()
	var filterModelList []FilterModel
	m := new(WhyInfo)
	err = json.Unmarshal([]byte(ParameterStr), &filterModelList)
	if err != nil {
		return nil, err
	}
	(*m).ColumnList = filterModelList
	return m.Reset(), err
}
func (m *WhyInfo) Reset() *WhyInfo {
	(*m).ParameterStr = ""
	(*m).Limt = *new(LimtModel)
	(*m).OrderByList = make([]OrderByModel, 0)
	(*m).TabName = ""
	(*m).CurrentPage = DefaultCurrentPage
	(*m).PageSize = DefaultPageSize
	(*m).IsLimt = false
	return m
}
func (m *WhyInfo) Clear() *WhyInfo {
	return m.Reset()
}
func (m *WhyInfo) SetTabName(tabName string) *WhyInfo {
	(*m).TabName = tabName
	return m
}
func (m *WhyInfo) SetPageSize(pageSize int) *WhyInfo {

	if pageSize <= 0 {
		pageSize = DefaultCurrentPage
	}
	(*m).PageSize = pageSize
	(*m).IsLimt = true
	return m
}
func (m *WhyInfo) SetCurrentPage(currentPage int) *WhyInfo {
	if currentPage <= 0 {
		currentPage = DefaultPageSize
	}
	(*m).CurrentPage = currentPage
	(*m).IsLimt = true
	return m
}
func (m *WhyInfo) SetLimt(args ...int) *WhyInfo {
	if len(args) <= 0 {
		(*m).CurrentPage = DefaultCurrentPage
		(*m).PageSize = DefaultPageSize
	} else if len(args) == 1 {
		temp := args[0]
		if temp <= 0 {
			temp = DefaultCurrentPage
		}
		(*m).CurrentPage = temp
		(*m).PageSize = DefaultPageSize
	} else if len(args) > 1 {
		temp := args[0]
		if temp <= 0 {
			temp = DefaultCurrentPage
		}
		temp1 := args[1]
		if temp1 <= 0 {
			temp1 = DefaultPageSize
		}
		(*m).CurrentPage = temp
		(*m).PageSize = temp1 - temp + 1
	}
	(*m).IsLimt = true
	return m
}
func (m *WhyInfo) SetOrderBy(rlist ...string) *WhyInfo {
	tlist := make([]OrderByModel, 0)
	if len(rlist) <= 0 {
		return m
	} else if len(rlist) == 1 {
		tlist = append(tlist, OrderByModel{Column: rlist[0], SortType: DESC})
	} else {
		if strings.ToUpper(rlist[len(rlist)-1]) != DESC && strings.ToUpper(rlist[len(rlist)-1]) != ASC {
			rlist = append(rlist, DESC)
		}
		for i := 0; i < len(rlist)-2; i++ {
			tlist = append(tlist, OrderByModel{Column: rlist[i], SortType: rlist[len(rlist)-1]})
		}
	}

	(*m).OrderByList = tlist
	return m
}
func (m *WhyInfo) SetOrderByCustomize(rlist ...OrderByModel) *WhyInfo {
	tlist := make([]OrderByModel, 0)
	for _, val := range rlist {
		tlist = append(tlist, val)
	}
	(*m).OrderByList = tlist
	return m
}
func (m *WhyInfo) getLimt() {
	PageSize := (*m).PageSize
	CurrentPage := (*m).CurrentPage
	min := (CurrentPage - 1) * PageSize
	max := CurrentPage * PageSize
	var temp LimtModel
	temp.Min = min
	temp.Max = max
	(*m).Limt = temp

}
func (m *WhyInfo) getLimtStr() {
	tempStr := ""
	(*m).getLimt()
	if (*m).Limt != (LimtModel{}) && (*m).IsLimt {
		tempStr = " " + strconv.Itoa((*m).Limt.Min) + "," + strconv.Itoa((*m).Limt.Max) + " "
	}
	(*m).LimtStr = tempStr
}
func (m *WhyInfo) getOrderByStr() {
	tempStr := ""
	if (*m).OrderByList == nil || len((*m).OrderByList) <= 0 {

	} else {
		for i, val := range (*m).OrderByList {
			if val != (OrderByModel{}) {
				SortType := "Desc"
				if val.SortType != "" {
					SortType = val.SortType
				}
				tempStr = tempStr + " " + val.Column + " " + SortType + " "
				if len((*m).OrderByList) != (i + 1) {
					tempStr = tempStr + ", "
				}
			}

		}

	}
	(*m).OrderByStr = tempStr
}
func (m *WhyInfo) getWhereSqlStr() error {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("生成WHERER条件异常")
		}
	}()

	tempStr := ""
	for i := 0; i < len((*m).ColumnList); i++ {
		fieldWhere, err := getFieldWhere(&((*m).ColumnList[i]))
		if err != nil {
			return err
		}
		if tempStr == "" {
			tempStr = "1=1 " + fieldWhere
		} else {
			tempStr += fieldWhere
		}

	}
	(*m).WhereStr = tempStr
	return err
}
func (m *WhyInfo) getWhereInit() error {

	tempStr := ""
	(*m).Str = tempStr
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("生成WHERER条件异常")
		}
	}()

	(*m).getOrderByStr()
	(*m).getLimtStr()
	err = (*m).getWhereSqlStr()
	if err != nil {
		return err
	}
	if (*m).WhereStr != "" {
		tempStr = " where " + (*m).WhereStr
	}

	if (*m).OrderByStr != "" {
		tempStr = tempStr + " OrderBy " + (*m).OrderByStr
	}

	if (*m).LimtStr != "" {
		tempStr = tempStr + " LIMIT " + (*m).LimtStr
	}
	(*m).Str = tempStr
	return err
}
func (m *WhyInfo) check() error {

	if (*m).TabName == "" {
		return errors.New("表名为空")
	}
	return nil

}
func (m *WhyInfo) GetQuerySql() (string, error) {
	tempStr := ""
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("生成SQL异常")
		}
	}()
	err = (*m).check()
	if err != nil {
		return "", err
	}
	err = (*m).getWhereInit()
	if err != nil {
		return "", err
	}
	tempStr = "select * from " + (*m).TabName + (*m).Str
	return tempStr, nil

}
func (m *WhyInfo) GetCountSql() (string, error) {

	tempStr := ""
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("生成SQL异常")
		}
	}()
	err = (*m).check()
	if err != nil {
		return "", err
	}
	err = (*m).getWhereInit()
	if err != nil {
		return "", err
	}
	tempStr = "select count(1) as Num from " + (*m).TabName + (*m).Str
	return tempStr, nil
}
