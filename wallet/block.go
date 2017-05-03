package wallet

type Block struct {
	Hash              string `json:"hash"`
	Confirmations     int `json:"confirmations"`
	Size              int `json:"size"`
	Height            int `json:"height"`
	Version           int `json:"version"`
	Merkleroot        string `json:"merkleroot"`
	Mint              float64 `json:"mint"`
	Time              int `json:"time"`
	Nonce             int `json:"nonce"`
	Bits              string `json:"bits"`
	Difficulty        float64 `json:"difficulty"`
	Blocktrust        string `json:"blocktrust"`
	Chaintrust        string `json:"chaintrust"`
	Previousblockhash string `json:"previousblockhash"`
	Flags             string `json:"flags"`
	Proofhash         string `json:"proofhash"`
	Entropybit        int `json:"entropybit"`
	Modifier          string `json:"modifier"`
	Modifierv2        string `json:"modifierv2"`
	Tx                []string `json:"tx"`
}
