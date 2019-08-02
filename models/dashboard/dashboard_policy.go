package models

import "gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models"

type Policy struct {
	models.Proto
	Content string `datastore:",noindex"`
}

func DummyPolicy(uid string) *Policy {
	return &Policy{
		Proto: models.Proto{
			UID: uid,
		},
	}
}
