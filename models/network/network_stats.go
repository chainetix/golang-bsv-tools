package models

func NewNetworkStats(netCfg *Network) (*NetworkStats, error) {

	return &NetworkStats{
		NetworkProto: NetworkProto{},
		Label: "My Network Label",
		BlockSeconds: 15,
		BlockHeight: 1001,
		Projects: 124,
		Public: true,
	}, nil
}

type NetworkStats struct {
	NetworkProto
	Label string `json:"label"`
	Nodes int `json:"nodes"`
	BlockHeight int `json:"blockHeight"`
	BlockSeconds int `json:"blockSeconds"`
	Projects int `json:"projects"`
	Total_TxCount int `json:"txCount"`
	Total_Currencies int `json:"currencies"`
	Total_Users int `json:"users"`
	Public bool `json:"public"`
}
