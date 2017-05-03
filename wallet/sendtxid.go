package wallet

type SendTxid struct {
	Result string `json:"result"`
	Error  interface{} `json:"error"`
	ID     int `json:"id"`
}
