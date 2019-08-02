package models

// Sent from App Engine to the cold node.
type SignatureRequest struct {
	Network
	TX string
	Array []*TxData
	Key string
	Args []string
}

type AddressKeyPair struct {
	Address string `json:"address"`
	PubKey string `json:"pubkey"`
	PrivKey string `json:"privkey"`
}

type Unspent struct {
	Txid string `json:"txid"`
	Vout int `json:"vout"`
}

type TxData struct {
	Txid string `json:"txid"`
	Vout int `json:"vout"`
	ScriptPubKey string `json:"scriptPubKey"`
}
