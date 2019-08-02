package models

type Address struct {
	WalletProto
	IsDefault bool `json:"isDefault"`
	Seed string `json:"-"`
	Addr string `json:"-"`
}

func (self *Address) GetUID() string {
	return self.UID
}

func (self *Address) Tree() []string {
	return []string{
		self.Project,
		self.User,
		self.Wallet,
	}
}

func (address *Address) NewBalance() *Balance {
	return &Balance{
		AddressProto: AddressProto{
			WalletProto: WalletProto{
				UserProto: UserProto{
					ProjectProto: ProjectProto{
						Project: address.Project,
					},
					User: address.User,
				},
				Wallet: address.Wallet,
			},
			Address: address.UID,
		},
	}
}

func (address *Address) NewPermission() *Permission {
	return &Permission{
		AddressProto: AddressProto{
			WalletProto: WalletProto{
				UserProto: UserProto{
					ProjectProto: ProjectProto{
						Project: address.Project,
					},
					User: address.User,
				},
				Wallet: address.Wallet,
			},
			Address: address.UID,
		},
	}
}
