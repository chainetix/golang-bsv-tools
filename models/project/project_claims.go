package models

import (
	"github.com/dgrijalva/jwt-go"
)

type ProjectClaims struct {
	jwt.StandardClaims
	Project string
	// uri pattern for controlling which resources permission has been granted for
	Resources []string
}

func (claims *ProjectClaims) Valid() error {
	return nil
}
