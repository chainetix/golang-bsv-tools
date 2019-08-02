package models

type Currency struct {
	ProjectProto
	Name string `json:"name"`
	Alias string `json:"-"`
	Units int `json:"units"`
	Open bool `json:"units"`
}

func (self *Currency) GetUID() string {
	return self.UID
}

func (self *Currency) Tree() []string {
	return []string{
		self.Project,
	}
}

func (currency *Currency) NewAsset() *Asset {
	return &Asset{
		CurrencyProto: CurrencyProto{
			ProjectProto: ProjectProto{
				Project: currency.Project,
			},
			Currency: currency.UID,
		},
	}
}
