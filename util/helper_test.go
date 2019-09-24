package util

import (
	"ahaschool.com/ahamkt/aha-go-library.git/cache/redis"
	"ahaschool.com/ahamkt/aha-go-library.git/container/pool"
	xtime "ahaschool.com/ahamkt/aha-go-library.git/time"
	"context"
	"encoding/json"
	"testing"
	"time"
)

var p *redis.Pool
var config *redis.Config

func init() {
	config = getConfig()
	p = redis.NewPool(config)
}

func getConfig() (c *redis.Config) {
	c = &redis.Config{
		Name:         "test",
		Proto:        "tcp",
		Addr:         "127.0.0.1:6379",
		DialTimeout:  xtime.Duration(time.Second),
		ReadTimeout:  xtime.Duration(time.Second),
		WriteTimeout: xtime.Duration(time.Second),
	}
	c.Config = &pool.Config{
		Active:      20,
		Idle:        2,
		IdleTimeout: xtime.Duration(90 * time.Second),
	}
	return
}

func TestGet(t *testing.T) {
	var (
		key  = "exchange_code_used_key"
		conn = p.Get(context.TODO())
	)
	defer conn.Close()
	var attrs map[string]string
	if reply, err := conn.Do("GET", key); err != nil {
		t.Errorf("redis: conn.Do(GET, %s) error(%v)", key, err)
	} else {

		if err := json.Unmarshal(reply.([]byte), &attrs); err != nil {
			t.Log(err)
		} else {
			t.Log(attrs)
		}
	}
	attrs["a"] = "b"
	mjson, _ := json.Marshal(attrs)
	if err := conn.Send("SET", key, string(mjson)); err != nil {
		t.Log(err)
	}

}

func TestA(t *testing.T) {
	rule := make(map[string]string)
	m := UniqueExchangeCode(rule, 8)
	t.Log(rule)
	t.Log(m)

	tt := aha_exchange_int2code(4481030252, 8)
	t.Log(tt)

}
