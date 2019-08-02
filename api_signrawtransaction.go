package bitcoinsv

import (
	"github.com/chainetix/bitcoinsvd-client/models/network"
)

func (client *Client) SignRawTransaction(rawTransaction string, txDataArray []*models.TxData, privateKey string, args ...string) (Response, error) {

	if txDataArray == nil {
		txDataArray = []*models.TxData{}
	}

	params := []interface{}{
		rawTransaction,
		txDataArray,
		[]string{privateKey},
	}
	for _, arg := range args {
		params = append(params, arg)
	}

	msg := client.Command(
		"signrawtransaction",
		params,
	)

	return client.Post(msg)
}
