package models

type Exchange struct {
	ProjectProto
	Expiry int64 `json:"expiry"`
	GiveCurrency string `json:"giveCurrency"`
	GiveQuantity float64 `json:"giveQuantity"`
	RecvCurrency string `json:"recvCurrency"`
	RecvQuantity float64 `json:"recvQuantity"`
	Tx string `json:"-"`
}

func (self *Exchange) GetUID() string {
	return self.UID
}

func (self *Exchange) Tree() []string {
	return []string{
		self.Project,
	}
}
