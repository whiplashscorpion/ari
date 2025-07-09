package ari

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

func (c *CommandClient) PlaybackGet(ctx context.Context, playbackId string) (Playback, error) {
	path, err := url.JoinPath("playbacks", playbackId)
	if err != nil {
		return Playback{}, fmt.Errorf("failed to build playback path: %w", err)
	}
	res, err := c.httpGet(ctx, path)
	if err != nil {
		return Playback{}, fmt.Errorf("failed to get playback: %w", err)
	}

	var output Playback
	err = json.Unmarshal(res, &output)
	if err != nil {
		return output, fmt.Errorf("failed to unmarshal playback: %w", err)
	}
	return output, nil
}

func (c *CommandClient) PlaybackStop(ctx context.Context, playbackId string) error {
	path, err := url.JoinPath("playbacks", playbackId)
	if err != nil {
		return fmt.Errorf("failed to build playback path: %w", err)
	}
	_, err = c.httpDelete(ctx, path)
	if err != nil {
		return fmt.Errorf("failed to stop playback: %w", err)
	}
	return nil
}

func (c *CommandClient) PlaybackControl(ctx context.Context, playbackId string, operation string) error {

	path, err := url.JoinPath("playbacks", playbackId)
	if err != nil {
		return fmt.Errorf("failed to build playback path: %w", err)
	}

	params := url.Values{}
	params.Set("operation", operation)
	path = path + "?" + params.Encode()

	_, err = c.httpPost(ctx, path, nil)
	if err != nil {
		return fmt.Errorf("failed to control playback: %w", err)
	}
	return nil
}
