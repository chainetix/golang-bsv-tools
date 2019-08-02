package models

type Permission struct {
	AddressProto
	Action string `json:"action"`
	State bool `json:"state"`
}
