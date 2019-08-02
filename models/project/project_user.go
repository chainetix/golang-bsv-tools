package models

type User struct {
	ProjectProto
	Agent bool `json:"agent"`
	Label string `json:"label"`
}

func (self *User) GetUID() string {
	return self.UID
}

func (self *User) Tree() []string {
	return []string{
		self.Project,
	}
}

func (user *User) NewWallet(defaultAddress *Address) *Wallet {
	return &Wallet{
		UserProto: UserProto{
			ProjectProto: ProjectProto{
				Project: user.Project,
			},
			User: user.UID,
		},
		DefaultAddress: defaultAddress.Addr,
	}
}

func (user *User) NewUserGroup() *UserGroup {
	return &UserGroup{
		UserProto: UserProto{
			ProjectProto: ProjectProto{
				Project: user.Project,
			},
			User: user.UID,
		},
	}
}
