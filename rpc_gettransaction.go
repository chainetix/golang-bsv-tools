package bitcoinsv

func (client *Client) GetTransaction(id string) (Response, error) {

	msg := client.Command(
		"gettransaction",
		[]interface{}{
			id,
		},
	)

	return client.Post(msg)
}
