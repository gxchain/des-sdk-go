package main

import (
	merchantClient "des-sdk-go/des-sdk-go"
)

func main() {
	client := merchantClient.NewMerchantClient("5K8iH1jMJxn8TKXXgHJHjkf8zGXsbVPvrCLvU2GekDh2nk4ZPSF", "1.2.323", "http://192.168.1.124:6388")
	params := map[string]interface{}{"bankCardNo": "6236681540015259109"}
	createExchangeResult := client.CreateDataExchangeRequest(9, params)
	client.GetResult(createExchangeResult, 8)

}