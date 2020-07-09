package whysql

type FilterModel struct {
	Column      string `json:"column"`      //字段名
	Action      string `json:"action"`      //操作符 > < =
	Logic       string `json:"logic"`       //关系  and or
	Value       string `json:"value"`       //值
	DataType    string `json:"dataType"`    //数据类型
	OrderByType string `json:"OrderByType"` //排序类型为空代表不参与排序

}
type OrderByModel struct {
	Column   string `json:"column"`   //字段名
	SortType string `json:"sortType"` //排序类型
}
type LimtModel struct {
	Min int
	Max int
}
type WhyInfo struct {
	BaseOrderByList []OrderByModel
	ColumnList      []FilterModel
	TabName         string
	ParameterStr    string
	OrderByList     []OrderByModel
	Limt            LimtModel
	PageSize        int
	CurrentPage     int
	LimtStr         string
	IsLimt          bool
	OrderByStr      string
	WhereStr        string
	Str             string
}
