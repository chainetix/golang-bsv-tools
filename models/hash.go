package models

import (
	"golang.org/x/crypto/sha3"
)

func Hash128(b []byte) []byte {
	digest := make([]byte, 16)
	sha3.ShakeSum128(digest, b)
	return digest
}

func Hash256(b []byte) []byte {
	digest := make([]byte, 32)
	sha3.ShakeSum256(digest, b)
	return digest
}
