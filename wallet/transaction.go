package wallet

type TransactionAccountRecord struct {
	Account       string `json:"account"`
	Address       string `json:"address"`
	Category      string `json:"category"`
	Amount        float64 `json:"amount"`
	Fee           float64 `json:"fee"`
	Confirmations int `json:"confirmations"`
	Blockhash     string `json:"blockhash"`
	Blockindex    int `json:"blockindex"`
	Blocktime     int `json:"blocktime"`
	Txid          string `json:"txid"`
	Time          int `json:"time"`
	Timereceived  int64 `json:"timereceived"`
}

type TransactionAccountRecords [] TransactionAccountRecord

func (a TransactionAccountRecords) Len() int { // 重写 Len() 方法
	return len(a)
}

func (a TransactionAccountRecords) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}

func (a TransactionAccountRecords) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].Timereceived < a[i].Timereceived
}

type Transactions struct {
	Result TransactionAccountRecords `json:"result"`
	Error  interface{} `json:"error"`
	ID     int `json:"id"`
}
