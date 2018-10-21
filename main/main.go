package main

import (
	merchantClient "des-sdk-go/des-sdk-go"
	"fmt"
)

func main() {
	client := merchantClient.NewMerchantClient("5K8iH1jMJxn8TKXXgHJHjkf8zGXsbVPvrCLvU2GekDh2nXXXXXX", "1.2.323", "http://192.168.1.124:XXXX")
	params := map[string]interface{}{"bankCardNo": "6236681540015XXXXXX"}
	createExchangeResult := client.CreateDataExchangeRequest(9, params)
	response := client.GetResult(createExchangeResult, 8)
	//fmt.Println(response)
	respArr, _ := response.Get("datasources").Array()
	if respArr != nil {
		for _, resp := range respArr {
			fmt.Println(resp)
		}
	}

}