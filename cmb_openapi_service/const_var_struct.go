package cmb_openapi_service

import (
	"github.com/tjfoc/gmsm/sm2"
	"sync"
)

const (
	CmbOpenApiBaseUrl = "https://api.cmbchina.com"
	verify            = "SM3withSM2"
	AppId             = "${替换为您在 Open API 中创建的应用的 appId}"     // todo replace with real value
	AppSecret         = "${替换为您在 Open API 中创建的应用的 appSecret}" // todo replace with real value
	PrivateKey        = "${替换为您的私钥}"                          // todo replace with real value
)

var (
	once          sync.Once
	sm2PrivateKey *sm2.PrivateKey
)

type (
	// CmbRemitDeclareInfo 汇出申报信息
	CmbRemitDeclareInfo struct {
		MerchNo          string  `json:"merchNo" name:"招行商户号（子单商户号）"`
		OrderNo          string  `json:"orderNo" name:"商户订单号（子单订单号）"`
		OrderDate        string  `json:"orderDate" name:"招行商户订单日期"` // yyyy-MM-dd 以招行订单日期为准
		MainOrderMerchNo string  `json:"mainOrderMerchNo" name:"母单招行商户号"`
		MainOrderNo      string  `json:"mainOrderNo" name:"母单商户订单号"`
		PayerName        string  `json:"payerName" name:"付款人名称"`
		PayerIdNo        string  `json:"payerIdNo" name:"付款人身份证号"`
		LogisticsDate    string  `json:"logisticsDate" name:"订单物流日期"` // yyyy-MM-dd
		OrderAmount      float64 `json:"orderAmount" name:"订单金额"`
		RemitFlag        string  `json:"remitFlag" name:"汇出标志"` //  Y：需要汇出
	}

	// CmbOpenApiRes 招行 openApi 接口响应
	CmbOpenApiRes struct {
		Code      string      `json:"code"`
		Message   string      `json:"message"`
		Data      interface{} `json:"data"`
		Timestamp string      `json:"timestamp"`
	}
)
