package sql

import "reflect"

func FormatSql(model interface{}, fields map[string]bool) string {
	t := reflect.TypeOf(model)
	fv := reflect.ValueOf(model).MethodByName("TableName")
	tableName := fv.Call(nil)[0].String()
	sql := "SELECT "
	for i := 0; i < t.Elem().NumField(); i++ {
		field := t.Elem().Field(i).Tag.Get("json")
		if _, ok := fields[field]; len(fields) > 0 && !ok {
			continue
		}
		if flag, ok := fields[field]; len(fields) > 0 && ok && !flag {
			continue
		}

		sql = sql + field + ","
	}
	sql = sql[:len(sql)-1] + " FROM " + tableName

	return sql
}
