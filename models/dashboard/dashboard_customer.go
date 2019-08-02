package models

import (
	//
	"firebase.google.com/go/auth"
	"github.com/golangdaddy/go.uuid"
)

type Customer struct {
	CusProto
	Email string `json:"-"`
	Challenge string `json:"-"`
	// time last privacy policy was accepted (so we know if to propmt user)
	AcceptedUpdate int64 `json:"acceptedUpdate"`
	// network permissions
	CanRead []string `json:"canRead"`
	CanWrite []string `json:"canWrite"`
	CanAdmin []string `json:"canAdmin"`
}

func DummyCustomer(uid string) *Customer {
	return &Customer{
		CusProto: CusProto{
			UID: uid,
		},
	}
}

func NewCustomer(token *auth.Token) *Customer {
	customer := &Customer{
		Email: token.Claims["email"].(string),

	}
	customer.UID = token.UID
	customer.NewChallenge()
	return customer
}

func (customer *Customer) NewChallenge() {
	challenge, _ := uuid.NewV4()
	customer.Challenge = challenge.String()
}
