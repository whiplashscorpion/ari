package ari

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

func (c *CommandClient) EndpointsList(ctx context.Context) ([]Endpoint, error) {
	path := "/endpoints"
	var output []Endpoint

	// process the request
	result, err := c.httpGet(ctx, path)
	if err != nil {
		return output, fmt.Errorf("failed to get endpoints: %w", err)

	}

	// unmarshall response body into struct
	err = json.Unmarshal(result, &output)
	if err != nil {
		return output, fmt.Errorf("failed to unmarshal endpoints: %w", err)
	}

	return output, nil

}

func (c *CommandClient) EndpointsSendMessage(ctx context.Context, to string, from string) error {
	path := "/endpoints/sendMessage"
	u, err := url.Parse(path)
	if err != nil {
		return fmt.Errorf("failed to parse endpoint message path: %w", err)
	}
	v := url.Values{}
	v.Add("to", to)
	v.Add("from", from)

	u.RawQuery = v.Encode()

	_, err = c.httpPut(ctx, u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to send endpoint message: %w", err)
	}

	return nil

}
