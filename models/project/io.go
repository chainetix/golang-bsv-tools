package models

import (
	"time"
)

type Input struct {
	TxId string `json:"txid"`
	Vout int `json:"vout"`
	Value float64 `json:"value,omitempty"`
	Created time.Time `json:"-"`
}

type Output struct {
	Asset string `json:"asset"`
	Amount float64 `json:"amount"`
	Created time.Time `json:"-"`
}

type AddressKeyPair struct {
	Address string `json:"address"`
	PubKey string `json:"pubkey"`
	PrivKey string `json:"privkey"`
}

type TxData struct {
	Txid string `json:"txid"`
	Vout int `json:"vout"`
	ScriptPubKey string `json:"scriptPubKey"`
}
