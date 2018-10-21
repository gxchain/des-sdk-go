Skip to content
 
Search or jump to…

Pull requests
Issues
Marketplace
Explore
 @hot3246624 Sign out
2
0 0 gxchain/des-sdk-py
 Code  Issues 0  Pull requests 0  Projects 0  Wiki  Insights  Settings
des-sdk-py/ 
README.md
  or cancel
    
 
1
# Install
2
go get "github.com/gxchain/des-sdk-go"
3
# Usage
4
## Merchant client
5
    //initializing
6
    client = MerchantClient('5K8iH1jMJxn8TKXXgHJHjkf8zGXsbVPvrCLvU2GekDh2nk4XXXX', accountID, url)
7
    //creat data exchange and return request id
8
    params := map[string]interface{}{"bankCardNo": "6236681540015XXXXXX"}
9
    result = client.createDataExchangeRequest(9, {    "bankCardNo": "bankID"
10
    //get result thourgh reuqest id
11
    response = client.GetResult(createExchangeResult, 8)
12
    respArr, _ := response.Get("datasources").Array()
13
    if respArr != nil {
14
        for _, resp := range respArr {
15
            fmt.Println(resp)
16
    else:
17
        fmt.Println("response is None, please try more")
18

22
​
23
# Dev Documents
24
https://doc.gxb.io/des/
25
​
@hot3246624
Commit changes

Update README.md

Add an optional extended description…
  Commit directly to the master branch.
  Create a new branch for this commit and start a pull request. Learn more about pull requests.
 
© 2018 GitHub, Inc.
Terms
Privacy
Security
Status
Help
Contact GitHub
Pricing
API
Training
Blog
About
Press h to open a hovercard with more details.
