package models

import "gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models"

func NewBillingRecord(namespace *Project, transaction *Transaction) *BillingRecord {
	return &BillingRecord{
		Project: namespace.UID,
		Transaction: *transaction,
	}
}

type BillingRecord struct {
	models.Proto
	Project string `json:"namespace"`
	Transaction
}
