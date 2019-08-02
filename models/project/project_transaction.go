package models

import (
	"time"
)

type TX struct {
	Id string
	Height int
	Hash string
}

func NewTransaction(srcUser, destUser *User, srcAddress, destAddress *Address, currency *Currency, quantity float64) *Transaction {

	transaction :=  &Transaction{
		AddressProto: AddressProto{
			WalletProto: WalletProto{
				UserProto: UserProto{
					ProjectProto: ProjectProto{
						Project: srcAddress.Project,
					},
					User: srcAddress.User,
				},
				Wallet: srcAddress.Wallet,
			},
			Address: srcAddress.UID,
		},
		Sender: srcUser.UID,
		Receiver: destUser.UID,
		SendAddress: srcAddress.Addr,
		SendAddressUID: srcAddress.UID,
		ReceiveAddress: destAddress.Addr,
		ReceiveAddressUID: destAddress.UID,
		OutputsTotal: quantity,
		Time: time.Now().UTC().Unix(),
	}
	if currency != nil {
		transaction.Currency = currency.Alias
	}
	return transaction
}

type Transaction struct {
	AddressProto
	TxId string `json:"txid"`
	Sender string `json:"sender"`
	Receiver string `json:"receiver"`
	SendAddress string `json:"sendAddress"`
	SendAddressUID string `json:"sendAddressUID"`
	ReceiveAddress string `json:"receiveAddress"`
	ReceiveAddressUID string `json:"receiveAddressUID"`
	Currency string `json:"currency"`
	OutputsTotal float64 `json:"outputsTotal"`
	Size int `json:"outputs"`
	Time int64 `json:"time"`
}

func (self *Transaction) GetUID() string {
	return self.UID
}

func (self *Transaction) Tree() []string {
	return []string{
		self.Project,
		self.User,
		self.Wallet,
		self.Address,
	}
}
