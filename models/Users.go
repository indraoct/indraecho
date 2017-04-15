package models

import (
	"indraecho/config"
)

type UserData struct {
	Err error
	DatabaseUser  string
	DatabasePassword string
}

func GetuserList() ([]map[string]interface{},error) {
	config.ConnectDB()
	rows,_ := config.Db.Query("select id, username, password from users")

	defer rows.Close()
	config.CloseDB()
	columns,_ := rows.Columns()

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}


	return tableData,nil
}



