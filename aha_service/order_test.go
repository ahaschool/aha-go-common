package aha_service

import (
	"testing"
)

func TestPostCreateOrder(t *testing.T) {

	//req := ReqOrder{
	//	Host: "http://dev:9898/hjm_order_war",
	//	Token: "",
	//}
	//
	//c, _ := gin.CreateTestContext(httptest.NewRecorder())
	//c.Header("Content-Type", "application/json; charset=utf-8")
	//c.Request, _ = http.NewRequest("POST", req.Host, nil)
	//ctx := http2.NewContext(c)
	//
	//orderType := 3
	//attributes := make(map[string]interface{})
	//attributes["product_id"] = 1234
	//attributes["option_id"] = 1234
	//attributes["user_id"] = 1000399
	//attributes["order_title"] = "测试订单"
	//attributes["pattern_id"] = 0
	//attributes["order_count"] = 1
	//attributes["apply_mode"] = 3
	//attributes["order_price"] = 0
	//attributes["order_status"] = 3
	//attributes["payment_status"] = 3
	//attributes["payment_from"] = 0
	//attributes["app_type"] = 3
	//attributes["isInApp"] = true
	//attributes["page_source"] = "product_detail"
	//attributes["address"] = map[string]int{"address_id":0}
	//rr, err := req.PostCreateOrder(ctx, attributes, orderType)
	//assert.Equal(t, nil, err)
	//
	//var rrObj map[string]interface{}
	//err = json.Unmarshal([]byte(rr), &rrObj)
	//assert.Equal(t, nil, err)
	//
	//assert.Equal(t, nil, err)
	//assert.Equal(t, float64(0), rrObj["code"])
}
