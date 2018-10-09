package des_sdk_go

import (
	//"net/http"
	"fmt"
	//"io/ioutil"
	//"github.com/bitly/go-simplejson"
	"reflect"
	"github.com/asmcos/requests"
	"github.com/bitly/go-simplejson"
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
	//resp, err  := http.Get(cli.BaseUrl + fmt.Sprintf("/api/product/%v", productId))
	//if err != nil {
	//	return nil, err
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println(body)
	//js, err := simplejson.NewJson(body)
	//fmt.Println(js)
	//result := js.Get("result")
	//return result, nil
	resp, err := requests.Get(cli.BaseUrl + fmt.Sprintf("/api/product/%v", productId))
	if err != nil {
		return nil, err
	}
	js, err := simplejson.NewJson(resp.Content())
	fmt.Println(js)

	//var json map[string]interface{}
	//resp.Json(&json)
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
	for _, datasourceAccount := range arr {
		fmt.Println(reflect.TypeOf(datasourceAccount))
		if account, ok := datasourceAccount.(map[string]interface{}); ok {
			fmt.Println("1111", account["accountId"])
		}
	}
	//data, ok := productResult["onlineDatasources"]
	//fmt.Println(data, reflect.TypeOf(data))
	//if !ok {
	//	panic("there is no online dataSources")
	//}
	//var dataExchangeReqList []interface{}
	//var req common.CreateDataExchangeResp

	//create request data
	//if items, ok := data.([]interface{}); ok {
	//	if len(items) == 0 {
	//		panic("there is no dataSources")
	//	}
	//	for _, item := range items{
	//		if lilItem, ok := item.(map[string]string); ok {
	//			if strings.Compare(cli.AccountId, lilItem["accountId"]) == 0{
	//				continue
	//			}
	//			req.RequestParams.Amount =
	//		}
	//	}
	//}
	//dataExchangeReqList = append(dataExchangeReqList, item)
	//fmt.Println(dataExchangeReqList, reflect.TypeOf(dataExchangeReqList))


	fmt.Println(productResult)
	fmt.Println(params)
	return 0
}