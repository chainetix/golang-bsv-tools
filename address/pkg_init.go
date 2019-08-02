package address

import (
	"encoding/hex"
)

var private_key_version []byte
var address_pubkeyhash_version []byte
var address_checksum_value []byte

func init() {

	config := Config{
		PrivateKeyVersion: "80",
		AddressPubkeyhashVersion: "00",
		AddressChecksumValue: "00000000",
	}

	private_key_version, _ = hex.DecodeString(config.PrivateKeyVersion)
	address_pubkeyhash_version, _ = hex.DecodeString(config.AddressPubkeyhashVersion)
	address_checksum_value, _ = hex.DecodeString(config.AddressChecksumValue)

}
