package models

import "gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models"

type ProjectProto struct {
	models.Proto
	Project string `json:"-"`
}

type CurrencyProto struct {
	ProjectProto
	Currency string `json:"-"`
}

type StreamProto struct {
	ProjectProto
	Stream string `json:"-"`
}

type UserProto struct {
	ProjectProto
	User string `json:"-"`
}

type WalletProto struct {
	UserProto
	Wallet string `json:"-"`
}

type AddressProto struct {
	WalletProto
	Address string `json:"-"`
}
