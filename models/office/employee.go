package models

import "gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models"

type Employee struct {
	models.Proto
	Email string
	Firstname string
	Lastname string
	Address_Line1 string
	Address_Line2 string
	Address_City string
	Address_PostCode string
	Bank_AccountName string
	Bank_AccountNumber string
	Bank_SortCode string
}

func DummyEmployee(uid string) *Employee {
	return &Employee{
		Proto: models.Proto{
			UID: uid,
		},
	}
}

func NewEmployee(email string) *Employee {
	return &Employee{
		Proto: models.Proto{},
		Email: email,
	}
}
