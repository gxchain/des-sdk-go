# Install
go get "github.com/gxchain/des-sdk-go"
# Usage
## Merchant client
      //initializing
      client = MerchantClient('5K8iH1jMJxn8TKXXgHJHjkf8zGXsbVPvrCLvU2GekDh2nk4XXXX', accountID, url)
   
      //creat data exchange and return request id
      params := map[string]interface{}{"bankCardNo": "6236681540015XXXXXX"}
   
      result = client.createDataExchangeRequest(9, {    "bankCardNo": "bankID"
   
      //get result thourgh reuqest id
      response = client.GetResult(createExchangeResult, 8)
   
      respArr, _ := response.Get("datasources").Array()
   
      if respArr != nil {
   
         for _, resp := range respArr {
       
              fmt.Println(resp)
      else:
   
         fmt.Println("response is None, please try more")
        
# Dev Documents
https://doc.gxb.io/des/
​
