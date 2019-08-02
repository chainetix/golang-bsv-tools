package bitcoinsv

func (client *Client) GetBlockchainInfo() (Response, error) {

	msg := client.Command(
		"getblockchaininfo",
		[]interface{}{},
	)

	return client.Post(msg)
}
