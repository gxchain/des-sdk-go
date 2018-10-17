package main

import (
	"crypto/elliptic"
	"math/big"
	"fmt"
)

func main() {
	//client := merchantClient.NewMerchantClient("5K8iH1jMJxn8TKXXgHJHjkf8zGXsbVPvrCLvU2GekDh2nk4ZPSF", "1.2.323", "http://192.168.1.124:6388")
	//params := map[string]interface{}{"bankCardNo": "6236681540015259109"}
	//client.CreateDataExchangeRequest(1, params)

	//x, y := elliptic.CurveParams{23,16,big.NewInt(5), 1,1}
	a,b := elliptic.P256().Params().Add(big.NewInt(1), big.NewInt(2), big.NewInt(1), big.NewInt(2))
	fmt.Println(a, b)
}