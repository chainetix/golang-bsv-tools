package models

type ApiToken struct {
	ProjectProto
	Token string `json:"token"`
	Digest string `json:"digest"`
	Resources string `json:"resources"`
}
