package ari

import (
	"context"
	"encoding/json"
	"net/url"
)

func (c *CommandClient) EndpointsList(ctx context.Context) ([]Endpoint, error) {
	path := "/endpoints"
	var output []Endpoint

	// process the request
	result, err := c.httpGet(ctx, path)
	if err != nil {
		return output, err

	}

	// unmarshall response body into struct
	err = json.Unmarshal(result, &output)
	if err != nil {
		return output, err
	}

	return output, nil

}

func (c *CommandClient) EndpointsSendMessage(ctx context.Context, to string, from string) error {
	path := "/endpoints/sendMessage"
	u, err := url.Parse(path)
	if err != nil {
		return err
	}
	v := url.Values{}
	v.Add("to", to)
	v.Add("from", from)

	u.RawQuery = v.Encode()

	_, err = c.httpPut(ctx, u.String(), nil)

	return err

}
