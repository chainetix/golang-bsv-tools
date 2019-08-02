package models

import (
	"fmt"
	"encoding/hex"
)

type Stream struct {
	ProjectProto
	PublicKey string
	Label string `json:"label"`
}

func (self *Stream) PubKey() ([]byte, error) {
	return hex.DecodeString(self.PublicKey)
}

func (self *Stream) GetUID() string {
	return self.UID
}

func (self *Stream) Tree() []string {
	return []string{
		self.Project,
	}
}

func (self *Stream) Key(key string) string {

	if len(key) == 0 {
		panic("INVALID KEY LENGTH FOR STREAM ITEM")
	}

	return fmt.Sprintf(
		"%s_%s",
		self.UID,
		key,
	)
}
