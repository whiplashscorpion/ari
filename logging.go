package ari

import (
	"context"
	"encoding/json"
	"fmt"
)

func (c *CommandClient) LoggingGetInfo(ctx context.Context) (LogChannel, error) {
	path := "/asterisk/logging"

	// process the request
	result, err := c.httpGet(ctx, path)
	if err != nil {
		return LogChannel{}, fmt.Errorf("failed to get logging info: %w", err)

	}

	var output LogChannel
	err = json.Unmarshal(result, &output)
	if err != nil {
		return output, fmt.Errorf("failed to unmarshal logging info: %w", err)
	}

	return output, nil

}
