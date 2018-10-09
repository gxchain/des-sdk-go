package common

type Amount struct {
	Amount uint `json:"amount" bson:"amount"`
	AssetId string `json:"assetId" bson:"assetId"`
}

type RequestParams struct {
	Memo string `json:"memo" bson:"memo"`
	Percent uint `json:"percent" bson:"percent"`
	Amount Amount `json:"amount" bson:"amount"`
	ProxyAccount string `json:"proxyAccount" bson:"proxyAccount"`
	Signatures []string `json:"signatures" bson:"signatures"`
	To string `json:"to" bson:"to"`
	From string `json:"from" bson:"from"`
	Expiration uint `json:"expiration" bson:"expiration"`
}

type CreateDataExchangeResp struct {
	Nonce int64 `json:"nonce" bson:"nonce"`
	Params string `json:"params" bson:"params"`
	RequestParams RequestParams `json:"requestParams" bson:"requestParams"`
}
