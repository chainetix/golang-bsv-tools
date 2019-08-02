package models

type NetworkProto struct {
	Model int `json:"-"`
	UID string `json:"uid"`
	Created int64 `json:"created"`
}
