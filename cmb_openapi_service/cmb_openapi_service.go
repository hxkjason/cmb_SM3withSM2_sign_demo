package cmb_openapi_service

import (
	"cmb_SM3withSM2_sign_demo/gmsm_service"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// GetSm2PrivateKeyInstance 获取sm2私钥实例
func GetSm2PrivateKeyInstance() *sm2.PrivateKey {

	once.Do(func() {
		sm2PrivateKey = gmsm_service.TransHexToSm2PrivateKey(PrivateKey)
	})

	return sm2PrivateKey
}

// OutRemitDeclare  汇出申报
func OutRemitDeclare(orderList []CmbRemitDeclareInfo) (CmbOpenApiRes, error) {

	params := map[string][]CmbRemitDeclareInfo{
		"orderList": orderList,
	}

	headers := map[string]string{
		"appid":        AppId,
		"verify":       verify,
		"Content-Type": "application/json;charset=utf-8",
		"funcCode":     "OVS_ACQ_DECLARE_INFORM",
		"sysCode":      "AP",
		"channel":      "AP",
	}

	requestUrl := CmbOpenApiBaseUrl + "/nms/ovsacq/acceptOrderDeclare"
	var response CmbOpenApiRes
	err := RequestCmbOpenApi(requestUrl, http.MethodPost, headers, params, &response)

	return response, err
}

// RequestCmbOpenApi 请求招行OpenApi
func RequestCmbOpenApi(requestUrl, method string, headers map[string]string, params, response interface{}) error {

	paramBytes, err := json.Marshal(params)
	if err != nil {
		return err
	}
	bodyStr := string(paramBytes)

	// 生成请求头sign (body 的 SM3 摘要)
	sign := gmsm_service.SM3Sum(bodyStr)
	timestampStr := strconv.FormatInt(time.Now().Unix(), 10)

	// 生成请求头 apiSign
	forSignStr := "appid=" + AppId + "&secret=" + AppSecret + "&sign=" + sign + "&timestamp=" + timestampStr
	apiSign, err := gmsm_service.SM3WithSM2Sign(GetSm2PrivateKeyInstance(), forSignStr, nil)
	if err != nil {
		return err
	}

	headers["timestamp"] = timestampStr
	headers["sign"] = sign
	headers["apisign"] = apiSign

	client := &http.Client{}
	req, _ := http.NewRequest(method, requestUrl, strings.NewReader(bodyStr))
	defer func() {
		_ = req.Body.Close()
	}()

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("CmbResponse:" + resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("resp:", string(body))
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &response)
}
