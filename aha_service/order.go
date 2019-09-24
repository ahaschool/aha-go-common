package aha_service

type ReqOrder struct {
	Host   string
	Token  string
}

// 创建订单
func (q *ReqOrder) PostCreateOrder(xenvStr string, attributes map[string]interface{}, orderType int) (response string, err error) {
	if attributes["address"] == nil {
		attributes["address"] = map[string]int{"address_id":0}
	}
	if attributes["page_source"] == nil {
		attributes["page_source"] = "product_detail"
	}
	if attributes["isInApp"] == nil {
		attributes["isInApp"] = true
	}
	if attributes["app_type"] == nil {
		attributes["app_type"] = 3
	}
	if attributes["payment_from"] == nil {
		attributes["payment_from"] = 0
	}
	if attributes["payment_status"] == nil {
		attributes["payment_status"] = 3
	}
	if attributes["order_status"] == nil {
		attributes["order_status"] = 3
	}
	if attributes["apply_mode"] == nil {
		attributes["apply_mode"] = 3
	}
	if attributes["order_count"] == nil {
		attributes["order_count"] = 1
	}
	if attributes["pattern_id"] == nil {
		attributes["pattern_id"] = 0
	}

	data := attributes
	if  orderType == 0 {
		orderType = 3
	}

	data["order_type"] = orderType
	conf := &Config{
		ReqHost: q.Host,
		XToken:  q.Token,
	}
	return ahaPost(conf, xenvStr, "/orders/system", data)
}
