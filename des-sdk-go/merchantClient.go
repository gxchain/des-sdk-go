package des_sdk_go

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/bitly/go-simplejson"
)

type MerchantClient struct {
	Client
}

func NewMerchantClient(privateKey string, account string, baseUrl string) *MerchantClient{
	cli := &MerchantClient{Client{privateKey,account,baseUrl}}
	return cli
}

func (cli *MerchantClient)getProduct(productId int) (*simplejson.Json, error) {
	resp, err  := http.Get(cli.BaseUrl + fmt.Sprintf("/api/product/%s", productId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	js, err := simplejson.NewJson(body)
	result := js.Get("result")
	return result, nil
}

func (cli *MerchantClient)CreateDataExchangeRequest(productId int, params map[string]string) int {
	productResult, err := cli.getProduct(productId)
	if err != nil {
		panic("get product result failed")
	}
	//if len(productResult.Get("onlineDatasources")) == 0 {
	//	panic("there is no online datasources")
	//}

}

func main() {
	client := NewMerchantClient("5K8iH1jMJxn8TKXXgHJHjkf8zGXsbVPvrCLvU2GekDh2nk4ZPSF", "1.2.323", "http://192.168.1.124:6388")

}