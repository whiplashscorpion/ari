package ari

import (
	"context"
	"encoding/json"
	"net/url"
)

// fairly happy with this setup, but haven't been able to test all scenarios

func (c *CommandClient) ConfigObjectGet(ctx context.Context, configClass string, objectType string, id string) (ConfigTuple, error) {
	path, err := url.JoinPath("/asterisk/config/dynamic", configClass, objectType, id)
	if err != nil {
		return ConfigTuple{}, err
	}

	var output ConfigTuple

	result, err := c.httpGet(ctx, path)
	if err != nil {
		return output, err
	}

	err = json.Unmarshal(result, &output)
	if err != nil {
		return output, err
	}

	return output, nil

}
func (c *CommandClient) ConfigObjectCreate(ctx context.Context, configClass string, objectType string, id string, fields ...[]ConfigTuple) (ConfigTuple, error) {
	path, err := url.JoinPath("/asterisk/config/dynamic", configClass, objectType, id)
	if err != nil {
		return ConfigTuple{}, err
	}

	// body
	body := []byte{}
	if fields != nil {
		body, err = json.Marshal(&fields[0])
		if err != nil {
			return ConfigTuple{}, err
		}
	}

	var output ConfigTuple

	result, err := c.httpPut(ctx, path, body)
	if err != nil {
		return output, err
	}

	err = json.Unmarshal(result, &output)
	if err != nil {
		return output, err
	}

	return output, nil

}
func (c *CommandClient) ConfigObjectDelete(ctx context.Context, configClass string, objectType string, id string) error {
	path, err := url.JoinPath("/asterisk/config/dynamic", configClass, objectType, id)
	if err != nil {
		return err
	}

	_, err = c.httpDelete(ctx, path)
	return err

}
