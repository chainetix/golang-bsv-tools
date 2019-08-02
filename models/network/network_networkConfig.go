package models

type Network struct {
	NetworkProto
	//
	Public bool `json:"public"`
	Shared bool `json:"shared"`
	//
	WizardKey string `json:"wizardKey"`
	RPCUser string `json:"rpcuser"`
	RPCPassword string `json:"rpcpassword"`
	//
	ChainName string `json:"chain-name"`
	DefaultRpcPort int `json:"default-rpc-port"`
	PrivateKeyVersion string `json:"private-key-version"`
	AddressPubkeyhashVersion string `json:"address-pubkeyhash-version"`
	AddressChecksumValue string `json:"address-checksum-value"`
	BurnAddress string `json:"burnaddress"`
}

func DummyNetwork(uid string) *Network {
	return &Network{
		NetworkProto: NetworkProto{
			UID: uid,
		},
	}
}

func toBool(s string) bool {
	return s == "true"
}
