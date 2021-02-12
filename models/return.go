package models

type ReturnJsonWithoutData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type PostReturnData struct {
	Id string `json:"id"`
}

type PostReturnJson struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data PostReturnData `json:"data"`
}

type DetailsReturnData struct {
	//JsonBody       map[string]interface{} `json:"json_body"`
	Url        string `json:"url"`
	Count      uint   `json:"count"`
	CreatedAt  string `json:"created_at"`
	LastUsedAt string `json:"last_used_at"`
	ExpiresAt  string `json:"expires_at"`
}

type DetailsReturnJson struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data DetailsReturnData `json:"data"`
}
