package models

type Child interface {
	GetUID() string
	Tree() []string
}
