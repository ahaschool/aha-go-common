package sql

import (
	"aha-api-server/src/srv/exchange/model"
	"testing"
)

func TestFormatSql(t *testing.T) {
	u := &model.ExchangeCode{}
	s1 := map[string]bool{
		"id":   false,
		"code": true,
	}
	res := FormatSql(u, nil)
	t.Log(res)
	res = FormatSql(u, s1)
	t.Log(res)
}
