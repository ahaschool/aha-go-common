package sql

import (
	"ahaschool.com/ahamkt/aha-go-library.git/net/netutil/breaker"
	xtime "ahaschool.com/ahamkt/aha-go-library.git/time"
	"aha-api-server/src/srv/exchange/model"
	"context"
	"fmt"
	"testing"
	"time"
)

func TestSql(t *testing.T) {
	bc := &breaker.Config{
		Window:  xtime.Duration(10 * time.Second),
		Sleep:   xtime.Duration(10 * time.Second),
		Bucket:  10,
		Ratio:   0.5,
		Request: 100,
	}
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "123456", "127.0.0.1", "3306", "operation")
	c := &Config{
		Addr:         "test",
		DSN:          dsn,
		Active:       10,
		Idle:         5,
		IdleTimeout:  xtime.Duration(time.Minute),
		QueryTimeout: xtime.Duration(time.Minute),
		ExecTimeout:  xtime.Duration(time.Minute),
		TranTimeout:  xtime.Duration(time.Minute),
		Breaker:      bc,
	}
	db := NewMySQL(c)
	var code model.ExchangeCode
	rows, err := db.Query(context.Background(), fmt.Sprintf("select code from h_exchange_code where code = 'AHAXZGGW'"))
	if err != nil {
		fmt.Println("====================")
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		r := new(model.ExchangeCode)
		if err = rows.Scan(&r.Code); err != nil {
			fmt.Println(err)
		}
		code = *r
	}
	fmt.Println(code.Code)
}
