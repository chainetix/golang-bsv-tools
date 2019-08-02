package models

import (
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models"
)

type StreamItem struct {
	models.Proto
	Blockhash     string   `json:"blockhash"`
	Blockindex    int      `json:"blockindex"`
	Blocktime     int      `json:"blocktime"`
	Confirmations int      `json:"confirmations"`
	Data          string   `json:"data"`
	Keys          []string `json:"keys"`
	Publishers    []string `json:"publishers"`
	Time          int      `json:"time"`
	Timereceived  int      `json:"timereceived"`
	Txid          string   `json:"txid"`
	Valid         bool     `json:"valid"`
	Vout          int      `json:"vout"`
}
