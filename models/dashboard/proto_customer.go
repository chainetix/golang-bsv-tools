package models

type CusProto struct {
	Model int `json:"-"`
	UID string `json:"uid"`
	Created int64 `json:"created"`
}

type CustomerProto struct {
	CusProto
	Customer string `json:"-"`
}
