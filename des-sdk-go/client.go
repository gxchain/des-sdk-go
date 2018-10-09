package des_sdk_go


type Client struct {
	PrivateKey string
	Account string
	BaseUrl string
}

func NewClient(privateKey string, account string, baseUrl string) *Client {
	client := &Client{privateKey,account, baseUrl}
	return client
}

