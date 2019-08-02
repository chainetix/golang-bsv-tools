package models

import (
	"github.com/golangdaddy/go.uuid"
)

type WalletItem struct {
	UID string `json:"uid"`
	Label string `json:"label"`
	Addresses []*Address `json:"addresses"`
}

type Wallet struct {
	UserProto
	Label string `json:"label"`
	DefaultAddress string `json:"defaultAddress"`
}

func (self *Wallet) GetUID() string {
	return self.UID
}

func (self *Wallet) Tree() []string {
	return []string{
		self.Project,
		self.User,
	}
}

func (wallet *Wallet) NewAddress(isDefault bool) (*Address, error) {

	address := &Address{
		WalletProto: WalletProto{
			UserProto: UserProto{
				ProjectProto: ProjectProto{
					Project: wallet.Project,
				},
				User: wallet.User,
			},
			Wallet: wallet.UID,
		},
	}

	if !isDefault {
		uid, err := uuid.NewV4()
		if err != nil {
			return nil, err
		}
		address.Seed = uid.String()
		address.IsDefault = isDefault
	}

	return address, nil
}
