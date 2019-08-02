package models

import (
	"encoding/json"
)

type FeedMessage struct {
	ProjectProto
	Target string `json:"target"`
	Type string `json:"type"`
	Subject string `json:"subject"`
	Parents string `json:"parents"`
}

func (self *FeedMessage) RawJSON() []byte {
	b, _ := json.Marshal(self)
	return b
}

func (self *FeedMessage) StringJSON() string {
	return string(self.RawJSON())
}
