package des_sdk_go

import (
	//"net/http"
	"fmt"
	//"io/ioutil"
	//"github.com/bitly/go-simplejson"
	"reflect"
	"github.com/asmcos/requests"
	"github.com/bitly/go-simplejson"
	"des-sdk-go/des-sdk-go/common"
	"time"
	"crypto/md5"
	"encoding/json"
	"math/rand"
	"strings"
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
	fmt.Println(js)
	return js, nil
}

func (cli *MerchantClient)CreateDataExchangeRequest(productId int, params map[string]interface{}) int {
	productResult, err := cli.getProduct(productId)
	fmt.Println(reflect.TypeOf(productResult))
	if err != nil {
		panic("get product result failed")
	}
	arr, err := productResult.Get("onlineDatasources").Array()
	if err != nil || len(arr) == 0 {
		panic("there is no datasource.")
	}
	fmt.Println("[][][]",arr)
	//TODO validate input params

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
		fmt.Println(reflect.TypeOf(datasourceAccount))
		if account, ok := datasourceAccount.(map[string]interface{}); ok {
			fmt.Println("1111", account["accountId"])
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
		createDataExchangeResp.RequestParams.Memo = fmt.Sprintf("%x", memo)
		createDataExchangeResp.RequestParams.Expiration = expiration
		requestParams := createDataExchangeResp.RequestParams
		createDataExchangeResp.RequestParams.Signatures = common.Signature(common.Serilization(requestParams), cli.PrivateKey) //TODO signature
		createDataExchangeResp.Nonce = rand.Int63()
		var publicKey string
		if account, ok := datasourceAccount.(map[string]interface{}); ok {
			fmt.Println("publicKey: ", account["publicKey"])
			publicKey = account["publicKey"].(string)
		}
		createDataExchangeResp.Params = common.Encrypt(cli.PrivateKey, publicKey, createDataExchangeResp.Nonce, param)
		dataExchangeReqList = append(dataExchangeReqList, createDataExchangeResp)
	}
	fmt.Println(productResult)
	fmt.Println(params)

	if len(dataExchangeReqList) == 0 {
		panic("dataExchange request is empty")
	}
	req := requests.Requests()
	req.Header.Set("Content-Type","application/json")
	resp, err := req.Post(cli.BaseUrl + fmt.Sprintf("/api/request/create/%d", productId), json.Marshal(dataExchangeReqList))
	if err != nil {
		panic("there is something wrong")
	}
	//if resp.R.StatusCode == 200 {
	//
	//}
	js, err := simplejson.NewJson(resp.Content())
	id, _ := js.Get("request_id").Int()
	return id
}

func (cli *MerchantClient)GetResult(requestId int, timeout int64) interface{} {
	start := time.Now().Unix()
	for {
		var a = 1
		resp, err := requests.Get(cli.BaseUrl + fmt.Sprintf("/api/request/%v", requestId))
		if err != nil {
			fmt.Println("there is something wrong")
		}
		dataExchange, err := simplejson.NewJson(resp.Content())
		status, _ := dataExchange.Get("status").String()
		if dataExchange != nil && strings.Compare(status, "IN_PROGRESS") == 0 {
			data, _ := dataExchange.Get("datasources").Array()
			if len(data) == 0 {
				return data
			}
			for _, dataExchangeDetail := range data {
				if data, ok := interface{}(dataExchangeDetail).(map[string]interface{}); ok {
					stauts, _ := data["status"].(string)
					if strings.Compare(stauts, "SUCCESS") != 0 {
						continue
					}
					data["data"] = cli.Decrypt(cli.PrivateKey,
						data["datasourcePublicKey"],
						data["nonce"],
						data["data"])
				}
			}
			return dataExchange
		}
		time.Sleep(0.06)
		if time.Now().Unix() - start > timeout {
			break
		}
	}
	return nil
}