package bitcoinsv

type Unspent struct {
	Txid string `json:"txid"`
	Vout int `json:"vout"`
}
