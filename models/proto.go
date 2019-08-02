package models

import (
	"time"
	//
	"github.com/golangdaddy/go.uuid"
)

type Proto struct {
	UID string `json:"uid"`
	Salt string `json:"-"`
	Created time.Time `json:"created"`
}

func NewProto() Proto {

	// generate unique uuids
	uid, _ := uuid.NewV4()
	salt, _ := uuid.NewV4()

	return Proto{
		UID: uid.String(),
		Salt: salt.String(),
		Created: time.Now().UTC(),
	}
}
