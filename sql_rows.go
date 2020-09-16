package helper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type SqlRows struct {

}

func NewSqlRows() *SqlRows {
	return &SqlRows{

	}
}

// GetRowsData 根据指定类型 获取 rows 数据列表
func (o *SqlRows) GetRowsData(rows *sql.Rows, srcMp map[string]interface{}) (listData []map[string]interface{}, err error) {
	columns, err := rows.Columns()
	if err != nil {
		return
	}

	NewSrcMap, err := o.DeepCopyJson(srcMp)
	if err != nil {
		return
	}

	listData = make([]map[string]interface{},0)
	for rows.Next() {
		row := make([]interface{}, len(columns))
		for i, f := range columns {
			var v interface{}
			v = NewSrcMap[f]
			row[i] = v
		}
		rowMap := make(map[string]interface{}, 0)
		for k, d := range row {
			rowMap[columns[k]] = d
		}
		listData = append(listData, rowMap)
	}
	return listData, nil
}

// GetRowsStringData 获取 rows 字符类型的数据列表
func (o *SqlRows) GetRowsStringData(rows *sql.Rows) (listData []map[string]interface{}, err error) {
	columns, err := rows.Columns()
	if err != nil {
		return
	}

	listData = make([]map[string]interface{},0)
	for rows.Next() {
		row := make([]interface{}, len(columns))
		for i, _ := range columns {
			var v interface{}
			row[i] = &v
		}
		err := rows.Scan(row...)
		if err != nil {

		}
		rowMap := make(map[string]interface{}, 0)
		for k, cv := range row {
			cs, ok := o.GetRowColumnStringValue(cv)
			if !ok {

			}
			rowMap[columns[k]] = cs
		}
		listData = append(listData, rowMap)
	}
	return listData, nil
}

// DeepCopyJson 使用 json 方式深拷贝数据
func (o *SqlRows)  DeepCopyJson(src map[string]interface{}) (dest map[string]interface{}, err error) {
	if src == nil {
		return
	}
	jsonStr, err := json.Marshal(src)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonStr, &dest)
	if err != nil {
		return
	}
	return
}

// GetRowColumnStringValue 获取数据库表记录 字段值是 string 类
func (o *SqlRows)  GetRowColumnStringValue(v interface{}) (str string, ok bool) {
	nv, ok := v.(*interface{})
	if ok {
		v = *nv
	}
	t := reflect.TypeOf(v)
	fmt.Println(t.Name(), t.Kind())
	switch v.(type) {
	case []uint8:
		bytes, _ := v.([]uint8)
		str = string(bytes)
		return
	case time.Time:
		t, _ := v.(time.Time)
		str = t.String()
	}
	return
}