package models

type Block struct {
	PrevHash int `json:"prev"`
	Hash int `json:"hash"`
	Height int `json:"height"`
	TxTotal int `json:"txTotal"`
}
