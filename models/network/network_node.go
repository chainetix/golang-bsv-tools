package models

import (
	"fmt"
	"strings"
	"strconv"
	"encoding/base64"
	//
	"github.com/golangdaddy/tarantula/web/validation"
)

type Node struct {
	NetworkProto
	IPv4 string `json:"ipv4"`
	RPCUser string `json:"rpcuser"`
	RPCPassword string `json:"-"`
}

func DummyNode(uid string) *Node {
	return &Node{
		NetworkProto: NetworkProto{
			UID: uid,
		},
	}
}

func NewNode(netCfg *Network, ipAddress, rpcuser, rpcpassword string) (*Node, error) {

	ipv4 := validation.IPv4{}
	for x, s := range strings.Split(ipAddress, ".") {
		i, err := strconv.Atoi(
			strings.TrimSpace(s),
		)
		if err != nil {
			return nil, err
		}
		ipv4[x] = i
	}

	return &Node{
		NetworkProto: NetworkProto{},
		IPv4: ipv4.String(),
		RPCUser: rpcuser,
		RPCPassword: rpcpassword,
	}, nil
}

func (node *Node) Credentials() string {
	return base64.StdEncoding.EncodeToString(
		[]byte(
			node.RPCUser + ":" + node.RPCPassword,
		),
	)
}

func (node *Node) Address() string {
	return fmt.Sprintf(
		"http://%s:%v",
		node.IPv4,
		4444,
	)
}
