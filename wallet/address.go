package wallet

type ValidateAddress struct {
	Isvalid      bool `json:"isvalid"`
	Address      string `json:"address"`
	Ismine       bool `json:"ismine"`
	Isscript     bool `json:"isscript"`
	Pubkey       string `json:"pubkey"`
	Iscompressed bool `json:"iscompressed"`
	Account      string `json:"account"`
}

