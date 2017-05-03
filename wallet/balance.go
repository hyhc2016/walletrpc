package wallet

type Balance struct {
	Result float64 `json:"result"`
	Error  interface{} `json:"error"`
	ID     int `json:"id"`
}

