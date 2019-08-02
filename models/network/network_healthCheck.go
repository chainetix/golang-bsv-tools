package models

import (
	"sync"
)

type HealthCheck struct {
	Seen int64 `json:"seen"`
	MessagesReceived int `json:"messagesReceived"`
	MessageErrors int `json:"messageErrors"`
	sync.RWMutex
}
