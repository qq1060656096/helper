package helper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type SqlRows struct {

}

func NewSqlRows() *SqlRows {
	return &SqlRows{

	}
}
func (o *SqlRows) GetRowsScanData(columns[]string, srcMp map[string]interface{}) (row []interface{}, err error) {
	if columns == nil {
		err = ErrRowsScanDataAssertTypeColumnsNil
		return
	}
	if srcMp == nil {
		err = ErrRowsScanDataAssertTypeSrcMpNil
		return
	}

	row = make([]interface{}, len(srcMp))
	for i, f := range columns {
		t := srcMp[f]
		switch t.(type) {
		case int64:
			var tt int64
			row[i] = &tt
		case float64:
			var tt float64
			row[i] = &tt
		case string:
			var tt string
			row[i] = &tt
		case bool:
			var tt bool
			row[i] = &tt
		default:
			err = fmt.Errorf("%w %T" ,ErrRowsScanDataAssertTypeNil, t)
			return
		}
	}
	return row, nil
}
// GetRowsData 根据指定类型 获取 rows 数据列表
func (o *SqlRows) GetRowsData(rows *sql.Rows, srcMp map[string]interface{}) (listData []map[string]interface{}, err error) {
	columns, err := rows.Columns()
	newSrcMp, err := o.DeepCopyJson(srcMp)
	if err != nil {
		return
	}
	listData = make([]map[string]interface{},0)
	for rows.Next() {
		row, err := o.GetRowsScanData(columns, newSrcMp)
		if err != nil {
			return nil, err
		}
		err = rows.Scan(row...)
		if err != nil {
			return nil, err
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
			return nil, err
		}
		rowMap := make(map[string]interface{}, 0)
		for k, cv := range row {
			cs, err := o.GetRowColumnStringValue(cv)
			if err != nil {
				return nil, err
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
	dest = make(map[string]interface{})
	err = json.Unmarshal(jsonStr, &dest)
	if err != nil {
		return nil, err
	}
	return
}

// GetRowColumnStringValue 获取数据库表记录 字段值是 string 类
func (o *SqlRows)  GetRowColumnStringValue(v interface{}) (str string, err error) {
	nv, ok := v.(*interface{})
	if ok {
		v = *nv
	}
	switch v.(type) {
		case []uint8:
			bytes, _ := v.([]uint8)
			str = string(bytes)
			return
		case time.Time:
			t, _ := v.(time.Time)
			str = t.String()
		default:
			err = fmt.Errorf("%w %T" ,ErrRowsStringDataAssertTypeNil, v)
			return
	}
	return
}