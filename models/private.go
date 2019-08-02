package models

import (
	"time"
)

type PrivateData struct {
	Key string
	Value string `datastore:",noindex"`
	Updated int64
}

func NewPrivateData(key string) *PrivateData {
	return &PrivateData{
		Key: key,
		Updated: time.Now().UTC().Unix(),
	}
}

func (self *PrivateData) Write(key []byte, s string) error {
	return nil
}
