package bitcoinsv

import (
	"fmt"
	"errors"
)

const (
	CONST_ID = "bitcoinsv-client"
)

type Response map[string]interface{}

func (r Response) Result() interface{} {
	return r["result"]
}

func (client *Client) IsDebugMode() bool {
	return client.debug
}

func (client *Client) DebugMode() *Client {
	client.debug = true
	return client
}

func (client *Client) msg(params []interface{}) map[string]interface{} {
	return map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"params": params,
	}
}

func (client *Client) Command(method string, params []interface{}) map[string]interface{} {

	msg := client.msg(params)
	msg["method"] = fmt.Sprintf("%s", method)

	return msg
}

func (client *Client) Post(msg interface{}) (Response, error) {

	fmt.Println("POSTING: "+client.host+" "+client.credentials, msg)

	obj := make(Response)
	_, err := client.http.Post(
		client.host,
		msg,
		&obj,
		map[string]string{
			"Authorization": "Basic " + client.credentials,
		},
	)
	if err != nil {
		return nil, err
	}

	if obj["error"] != nil {
		e := obj["error"].(map[string]interface{})
		var s string
		m, ok := msg.(map[string]interface{})
		if ok {
			s = fmt.Sprintf("bitcoinsvd - '%s': %s", m["method"], e["message"].(string))
		} else {
			s = fmt.Sprintf("bitcoinsvd - %s", e["message"].(string))
		}
		return nil, errors.New(s)
	}

	return obj, nil
}
