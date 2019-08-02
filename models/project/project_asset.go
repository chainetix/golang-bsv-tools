package models

import "gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models"

type Asset struct {
	CurrencyProto
	Recipient string `json:"recipient,omitempty"`
	Quantity float64 `json:"quantity"`
}

func DummyAsset(uid string) *Asset {
	return &Asset{
		CurrencyProto: CurrencyProto{
			ProjectProto: ProjectProto{
				Proto: models.Proto{
					UID: uid,
				},
			},
		},
	}
}

func (self *Asset) GetUID() string {
	return self.UID
}

func (self *Asset) Tree() []string {
	return []string{
		self.Project,
		self.Currency,
		self.UID,
	}
}
