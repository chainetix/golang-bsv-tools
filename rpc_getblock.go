package bitcoinsv

func (client *Client) GetBlock(hash string, verbosity bool) (Response, error) {

	msg := client.Command(
		"getblock",
		[]interface{}{
			hash,
			verbosity,
		},
	)

	return client.Post(msg)
}
