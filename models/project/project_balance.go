package models

type Balance struct {
	AddressProto
	Asset string `json:"asset"`
	Value float64 `json:"value"`
}
