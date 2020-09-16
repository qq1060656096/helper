# helper
Golang Helper


### 从数据库获取任意表记录
```go
package main

import 	(
    "github.com/qq1060656096/helper"
	"database/sql"
    "fmt"
)
func main() {
    m := map[string]interface{}{                             
        "id": 0,
        "name": "",
    }
    var rows sql.Rows
    listData, err := helper.NewSqlRows().GetRowsStringData(rows)
    listData, err = helper.NewSqlRows().GetRowsData(rows, m)
    fmt.Println(listData, err)
}

```