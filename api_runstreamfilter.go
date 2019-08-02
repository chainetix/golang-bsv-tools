package bitcoinsv

func (client *Client) RunStreamFilter(filter string) (Response, error) {

	msg := client.Command(
		"runstreamfilter",
		[]interface{}{
            filter,
		},
	)

	return client.Post(msg)
}
