package main

import (
	"cmb_SM3withSM2_sign_demo/cmb_openapi_service"
	"fmt"
)

// go run main.go
func main() {
	RequestRemitDeclareDemo()
}

// RequestRemitDeclareDemo 请求汇出申报Demo
func RequestRemitDeclareDemo() {

	orderList := []cmb_openapi_service.CmbRemitDeclareInfo{
		{
			MainOrderMerchNo: "308999170120GK3",          // 招行母单商户号
			MerchNo:          "308999160120006",          // 招行子单商户号
			MainOrderNo:      "2358327362443478",         // 商户母单订单号
			OrderNo:          "200224204630010020000486", // 商户子单订单号
			OrderDate:        "2021-03-16",               // 请以招行订单日期为准
			PayerName:        "ZhangSan",                 // 付款人名称
			PayerIdNo:        "352203123456780001",       // 付款人身份证号
			LogisticsDate:    "2021-03-17",               // 订单物流日期
			OrderAmount:      18.01,                      // 订单金额
			RemitFlag:        "Y",                        // 汇出标志
		},
	}

	resp, err := cmb_openapi_service.OutRemitDeclare(orderList)
	if err != nil {
		fmt.Println("HasErr:", err.Error())
		return
	}

	fmt.Println("code:", resp.Code)
	fmt.Println("message:", resp.Message)
	fmt.Println("data:", resp.Data)
	fmt.Println("timestamp:", resp.Timestamp)
}
