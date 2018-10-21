package des_sdk_go

import (
	"fmt"

	"github.com/asmcos/requests"
	"github.com/bitly/go-simplejson"
	"des-sdk-go/des-sdk-go/common"
	"time"
	"crypto/md5"
	"encoding/json"
	"math/rand"
	"strings"
	"net/http"
	"bytes"
	"io/ioutil"
)

type MerchantClient struct {
	PrivateKey string
	AccountId string
	BaseUrl string
}

func NewMerchantClient(privateKey string, account string, baseUrl string) *MerchantClient{
	cli := &MerchantClient{privateKey,account,baseUrl}
	return cli
}

func (cli *MerchantClient)getProduct(productId int) (*simplejson.Json, error) {
	resp, err := requests.Get(cli.BaseUrl + fmt.Sprintf("/api/product/%v", productId))
	if err != nil {
		return nil, err
	}
	js, err := simplejson.NewJson(resp.Content())
	//fmt.Println(js)
	return js, nil
}

func (cli *MerchantClient)CreateDataExchangeRequest(productId int, params map[string]interface{}) string {
	productResult, err := cli.getProduct(productId)
	//fmt.Println(reflect.TypeOf(productResult))
	if err != nil {
		panic("get product result failed")
	}
	arr, err := productResult.Get("onlineDatasources").Array()
	if err != nil || len(arr) == 0 {
		panic("there is no datasource.")
	}
	//fmt.Println("[][][]",arr)
	//TODO validate input params
	if common.VarifyParameters(params) {
		
	}else {
		
	}

	var createDataExchangeResp common.CreateDataExchangeResp
	var dataExchangeReqList []common.CreateDataExchangeResp
	param := map[string]interface{}{"params": params, "timestamp": time.Now().Unix()}
	expiration := time.Now().Unix() + common.DEFAULT_TIMEOUT
	createDataExchangeResp.RequestParams.Amount.Amount, err = productResult.Get("product").Get("price").Get("amount").Int()
	if err != nil {
		panic("there is no amount.")
	}
	createDataExchangeResp.RequestParams.Amount.AssetId, err = productResult.Get("product").Get("price").Get("assetId").String()
	if err != nil {
		panic("there is no assetId.")
	}
	for _, datasourceAccount := range arr {
		//fmt.Println(reflect.TypeOf(datasourceAccount))
		if account, ok := datasourceAccount.(map[string]interface{}); ok {
			createDataExchangeResp.RequestParams.To, ok = account["accountId"].(string)
		}
		createDataExchangeResp.RequestParams.From = cli.AccountId
		createDataExchangeResp.RequestParams.ProxyAccount, err = productResult.Get("des").Get("accountId").String()
		createDataExchangeResp.RequestParams.Percent, err = productResult.Get("des").Get("percent").Int()
		tempData, err := json.Marshal(param)
		if err != nil {
			fmt.Println("json.Marshal failed:", err)
			panic("json.Marshal failed")
		}
		memo := md5.Sum(tempData)
		createDataExchangeResp.RequestParams.Memo = fmt.Sprintf("memo %x", memo)
		createDataExchangeResp.RequestParams.Expiration = expiration
		requestParams := &createDataExchangeResp.RequestParams

		//create signature
		signature, err := common.Sign(requestParams, []string{cli.PrivateKey})
		createDataExchangeResp.RequestParams.Signatures = signature
		createDataExchangeResp.Nonce = uint64(rand.Int63())


		var publicKey string
		if account, ok := datasourceAccount.(map[string]interface{}); ok {
			//fmt.Println("publicKey: ", account["publicKey"])
			publicKey = account["publicKey"].(string)
		}
		//fmt.Println("parameter: ", tempData)
		if err != nil {
			panic("json.Marshal failed:")
		}
		//fmt.Println("encrypt:", cli.PrivateKey, publicKey, createDataExchangeResp.Nonce, tempData, string(tempData)) for test

		createDataExchangeResp.Params = common.Encrypt(cli.PrivateKey, publicKey, createDataExchangeResp.Nonce, tempData)
		dataExchangeReqList = append(dataExchangeReqList, createDataExchangeResp)
	}

	//fmt.Println(productResult)
	//fmt.Println(params)

	if len(dataExchangeReqList) == 0 {
		panic("dataExchange request is empty")
	}

	reqParam, err := json.Marshal(dataExchangeReqList)
	req, err := http.NewRequest("POST", cli.BaseUrl + fmt.Sprintf("/api/request/create/%d", productId), bytes.NewBuffer(reqParam))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("there is something wrong in http")
		return ""
	}
	defer resp.Body.Close()

	statuscode := resp.StatusCode
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	//fmt.Println(statuscode)
	//fmt.Println(statuscode, body, resp.Body, resp.Header)
	if err != nil {
		panic("there is something wrong")
	}
	resultId := ""
	if statuscode == 200 {

		js, _ := simplejson.NewJson(body)
		//fmt.Println(js)
		resultId, _ = js.Get("request_id").String()
	}

	return resultId
}

func (cli *MerchantClient)GetResult(requestId string, timeout int64) *simplejson.Json {
	start := time.Now().Unix()
	if requestId == "" {
		var result []string
		result = append(result, "create data exchange is something wrong.")
		resultAll := map[string]interface{}{"datasources": result}
		result2, _ := json.Marshal(resultAll)
		result3, _ := simplejson.NewJson(result2)
		return result3
	}
	for {
		resp, err := requests.Get(cli.BaseUrl + fmt.Sprintf("/api/request/%v", requestId))
		if err != nil {
			fmt.Println("there is something wrong")
		}
		dataExchange, err := simplejson.NewJson(resp.Content())
		status, _ := dataExchange.Get("status").String()
		fmt.Println(status)
		if strings.Compare(status, "FAIL") == 0 {
			return dataExchange
		}
		if dataExchange != nil && strings.Compare(status, "IN_PROGRESS") != 0 {
			data, _ := dataExchange.Get("datasources").Array()
			if len(data) == 0 {
				return dataExchange
			}
			for _, dataExchangeDetail := range data {
				if data, ok := interface{}(dataExchangeDetail).(map[string]interface{}); ok {
					stauts, _ := data["status"].(string)
					if strings.Compare(stauts, "SUCCESS") != 0 {
						continue
					}
					datasourcePublicKey, _ := data["datasourcePublicKey"].(string)
					nonce, _ := data["nonce"].(json.Number)
					data, _ := data["data"].(string)
					nonceNum, _ := nonce.Int64()
					data = common.Decrypt(cli.PrivateKey,
						datasourcePublicKey,
						nonceNum,
						data,
					)
					var result []string
					result = append(result,data)
					resultAll := map[string]interface{}{"datasources": result}
					result2, _ := json.Marshal(resultAll)
					result3, _ := simplejson.NewJson(result2)
					dataExchange = result3
				}
			}
			return dataExchange
		}
		time.Sleep(1)
		if time.Now().Unix() - start > timeout {
			break
		}
	}
	return nil
}