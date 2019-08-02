package models

import "gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models"

type Filter struct {
	models.Proto
	Name string `json:"name"`
	Type string `json:"type"`
	Restrictions string `json:"restrictions"`
	Code string `json:"code"`
	//Description string `json:"desc"`
}

func (self *Filter) GetUID() string {
	return self.UID
}

func NewFilter() *Filter {
	return &Filter{
		Proto: models.Proto{},
	}
}
