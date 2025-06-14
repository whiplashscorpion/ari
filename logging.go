package ari

import (
	"context"
	"encoding/json"
)

func (c *CommandClient) LoggingGetInfo(ctx context.Context) (*LogChannel, error) {
	path := "/asterisk/logging"

	// process the request
	result, err := c.httpGet(ctx, path)
	if err != nil {
		return nil, err

	}

	output := new(LogChannel)
	err = json.Unmarshal(result, output)
	if err != nil {
		return output, err
	}

	return output, nil

}
