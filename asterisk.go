package ari

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"slices"
)

func (c *CommandClient) AsteriskInfo(ctx context.Context, only ...string) (AsteriskInfo, error) {
	var output AsteriskInfo
	path := "/asterisk/info"

	// "only" filter is kind of useless. Return type AsteriskInfo is a struct
	// that groups all of the filtered structs anyway.
	// Included for sake of completeness

	if only != nil {
		valid := []string{"build", "system", "config", "status"}
		if slices.Contains(valid, only[0]) {
			qparams := url.Values{}
			qparams.Add("only", only[0])
			path = path + "?" + qparams.Encode()
		}
	}

	result, err := c.httpGet(ctx, path)
	if err != nil {
		return output, fmt.Errorf("failed to get asterisk info: %w", err)
	}

	err = json.Unmarshal(result, &output)
	if err != nil {
		return output, fmt.Errorf("failed to unmarshal asterisk info: %w", err)
	}

	return output, nil

}

func (c *CommandClient) AsteriskPing(ctx context.Context) (AsteriskPing, error) {
	path := "/asterisk/ping"

	// set the return type
	var output AsteriskPing

	// process the request
	result, err := c.httpGet(ctx, path)
	if err != nil {
		return output, fmt.Errorf("failed to ping asterisk: %w", err)

	}

	// unmarshall response body into struct
	err = json.Unmarshal(result, &output)
	if err != nil {
		return output, fmt.Errorf("failed to unmarshal asterisk ping: %w", err)
	}

	return output, nil

}

func (c *CommandClient) AsteriskVariableGet(ctx context.Context, variable string) (Variable, error) {
	var output Variable

	path := "/asterisk/variable"

	qparams := url.Values{}
	qparams.Add("variable", variable)
	path = path + "?" + qparams.Encode()

	result, err := c.httpGet(ctx, path)
	if err != nil {
		return output, err
	}

	err = json.Unmarshal(result, &output)
	if err != nil {
		return output, err
	}

	return output, err
}

func (c *CommandClient) AsteriskVariableSet(ctx context.Context, variable string, value string) error {
	path := "/asterisk/variable"
	qparams := url.Values{}
	qparams.Add("variable", variable)
	qparams.Add("value", value)
	path = path + "?" + qparams.Encode()

	_, err := c.httpGet(ctx, path)
	if err != nil {
		return err
	}

	return nil

}
