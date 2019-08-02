package models

type UserGroup struct {
	UserProto
	Name string `json:"groups"`
}
