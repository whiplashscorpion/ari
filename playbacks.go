package ari

import (
	"context"
	"encoding/json"
	"net/url"
)

func (c *CommandClient) PlaybackGet(ctx context.Context, playbackId string) (Playback, error) {
	path, err := url.JoinPath("playbacks", playbackId)
	if err != nil {
		return Playback{}, err
	}
	res, err := c.httpGet(ctx, path)
	if err != nil {
		return Playback{}, err
	}

	var output Playback
	err = json.Unmarshal(res, &output)
	return output, err
}

func (c *CommandClient) PlaybackStop(ctx context.Context, playbackId string) error {
	path, err := url.JoinPath("playbacks", playbackId)
	if err != nil {
		return err
	}
	_, err = c.httpDelete(ctx, path)
	return err
}

func (c *CommandClient) PlaybackControl(ctx context.Context, playbackId string, operation string) error {

	path, err := url.JoinPath("playbacks", playbackId)
	if err != nil {
		return err
	}

	params := url.Values{}
	params.Set("operation", operation)
	path = path + "?" + params.Encode()

	_, err = c.httpPost(ctx, path, nil)
	return err
}
