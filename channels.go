package ari

import (
	"encoding/json"

	"context"
	"fmt"
	"net/url"
	"slices"
	"strconv"
	"strings"
)

func (c *CommandClient) ChannelAnswer(ctx context.Context, channelId string) error {
	path, err := url.JoinPath("/channels", channelId, "answer")
	if err != nil {
		return err
	}

	_, err = c.httpPost(ctx, path, nil)

	return err

}

func (c *CommandClient) ChannelMute(ctx context.Context, channelId string, direction string) error {
	// direction: both, in, out

	path, err := url.JoinPath("channels", channelId, "mute")
	if err != nil {
		return err
	}

	accepted := []string{"both", "in", "out"}
	if !slices.Contains(accepted, direction) {
		return fmt.Errorf("invalid mute direction: %s", direction)
	}

	params := url.Values{}
	params.Set("direction", direction)
	path = path + "?" + params.Encode()

	_, err = c.httpPost(ctx, path, nil)
	return err
}

func (c *CommandClient) ChannelUnmute(ctx context.Context, channelId string, direction string) error {
	// direction: both, in, out

	path, err := url.JoinPath("channels", channelId, "mute")
	if err != nil {
		return err
	}

	// request will fail without direction query
	// this case will also catch empty string
	accepted := []string{"both", "in", "out"}
	if !slices.Contains(accepted, direction) {
		direction = "both"
	}

	params := url.Values{}
	params.Set("direction", direction)
	path = path + "?" + params.Encode()

	_, err = c.httpDelete(ctx, path)
	return err
}

func (c *CommandClient) ChannelList(ctx context.Context) ([]Channel, error) {
	path := "channels"

	result, err := c.httpGet(ctx, path)
	if err != nil {
		return []Channel{}, err
	}

	var output []Channel
	err = json.Unmarshal(result, &output)
	return output, err
}

func (c *CommandClient) ChannelOriginate(ctx context.Context, endpoint string, opts ChannelOpts) (Channel, error) {
	path := "/channels"

	params := url.Values{}
	params.Set("endpoint", endpoint)
	params.Set("app", opts.App)
	path = path + "?" + params.Encode()

	// override required fields
	opts.Endpoint = endpoint
	/*
		body, err := json.Marshal(opts)
		if err != nil {
			return Channel{}, err
		}*/

	result, err := c.httpPost(ctx, path, nil)
	if err != nil {
		return Channel{}, err
	}

	var output Channel
	err = json.Unmarshal(result, &output)
	return output, err
}
func (c *CommandClient) ChannelOriginateWithId(ctx context.Context, endpoint string, channelId string, opts ChannelOpts) (Channel, error) {
	path, err := url.JoinPath("/channels", channelId)
	if err != nil {
		return Channel{}, err
	}

	// override required fields
	opts.ChannelId = "" // empty because it's a path param
	opts.Endpoint = endpoint

	body, err := json.Marshal(opts)
	if err != nil {
		return Channel{}, err
	}

	result, err := c.httpPost(ctx, path, body)
	if err != nil {
		return Channel{}, err
	}

	var output Channel
	err = json.Unmarshal(result, &output)
	return output, err
}

func (c *CommandClient) ChannelCreate(ctx context.Context, endpoint string, app string, opts ChannelOpts) (Channel, error) {
	path := "/channels/create"

	// override required fields
	opts.Endpoint = endpoint
	opts.App = app

	body, err := json.Marshal(opts)
	if err != nil {
		return Channel{}, err
	}

	result, err := c.httpPost(ctx, path, body)
	if err != nil {
		return Channel{}, err
	}

	var output Channel
	err = json.Unmarshal(result, &output)
	return output, err
}

func (c *CommandClient) ChannelMove(ctx context.Context, channelId string, app string) error {
	path, err := url.JoinPath("channels", channelId, "move")
	if err != nil {
		return err
	}
	params := url.Values{}
	params.Set("app", app)

	path = path + "?" + params.Encode()

	_, err = c.httpPost(ctx, path, nil)
	return err
}

func (c *CommandClient) ChannelGet(ctx context.Context, channelId string) (Channel, error) {
	path, err := url.JoinPath("/channels", channelId)
	if err != nil {
		return Channel{}, err
	}

	result, err := c.httpGet(ctx, path)
	if err != nil {
		return Channel{}, err
	}

	var output Channel
	err = json.Unmarshal(result, &output)
	return output, err
}

func (c *CommandClient) ChannelHangup(ctx context.Context, channelId string, reason string) error {

	path, err := url.JoinPath("/channels", channelId)
	if err != nil {
		return err
	}

	if reason == "" {
		reason = "normal"
		// The reason field is optional in asterisk, so we use "normal".
		// "reason" is used over "reason_code" as reason_code is more linked to asterisk internal processes.
	} else {
		allowed := []string{
			"normal",
			"busy",
			"congestion",
			"no_answer",
			"timeout",
			"rejected",
			"unallocated",
			"normal_unspecified",
			"number_incomplete",
			"codec_mismatch",
			"interworking",
			"failure",
			"answered_elsewhere",
		}

		if !slices.Contains(allowed, reason) {
			return fmt.Errorf("invalid hangup reason: %s", reason)
		}
	}

	params := url.Values{}
	params.Set("reason", reason)
	path = path + "?" + params.Encode()

	_, err = c.httpDelete(ctx, path)
	return err

}

func (c *CommandClient) ChannelRecord(ctx context.Context, channelId string, filename string, format string, opts ...ChannelRecordOpts) (LiveRecording, error) {
	path, err := url.JoinPath("/channels", channelId, "record")
	if err != nil {
		return LiveRecording{}, err
	}

	params := url.Values{}
	params.Set("name", filename)
	params.Set("format", format)

	if len(opts) > 0 {
		options := opts[0]
		params.Set("maxDurationSeconds", strconv.Itoa(options.MaxDurationSeconds))
		params.Set("maxSilenceSeconds", strconv.Itoa(options.MaxSilenceSeconds))
		params.Set("ifExists", options.IfExists)
		params.Set("beep", strconv.FormatBool(options.Beep))
		params.Set("terminateOn", options.TerminateOn)
	}

	path = path + "?" + params.Encode()

	b, err := c.httpPost(ctx, path, nil)
	if err != nil {
		return LiveRecording{}, err
	}

	var output LiveRecording
	err = json.Unmarshal(b, &output)
	return output, err

}
func (c *CommandClient) ChannelRing(ctx context.Context, channelId string) error {
	path, err := url.JoinPath("/channels", channelId, "ring")
	if err != nil {
		return err
	}

	_, err = c.httpPost(ctx, path, nil)
	return err

}

func (c *CommandClient) ChannelRingStop(ctx context.Context, channelId string) error {
	path, err := url.JoinPath("/channels", channelId, "ring")
	if err != nil {
		return err
	}

	_, err = c.httpDelete(ctx, path)
	return err

}

func (c *CommandClient) channelPlay(ctx context.Context, pathBase, media string) (Playback, error) {
	params := url.Values{}
	params.Set("media", media)

	fullPath := pathBase + "?" + params.Encode()

	res, err := c.httpPost(ctx, fullPath, nil)
	if err != nil {
		return Playback{}, err
	}

	var output Playback
	err = json.Unmarshal(res, &output)
	return output, err
}

func (c *CommandClient) ChannelPlay(ctx context.Context, channelId string, media string) (Playback, error) {
	path, err := url.JoinPath("channels", channelId, "play")
	if err != nil {
		return Playback{}, err
	}
	return c.channelPlay(ctx, path, media)
}

func (c *CommandClient) ChannelPlayWithId(ctx context.Context, channelId string, media string, playbackId string) (Playback, error) {
	path, err := url.JoinPath("channels", channelId, "play", playbackId)
	if err != nil {
		return Playback{}, err
	}
	return c.channelPlay(ctx, path, media)
}

type ChannelOpts struct {
	App            string `json:"app,omitempty"`
	AppArgs        string `json:"appArgs,omitempty"`
	CallerId       string `json:"callerId,omitempty"`
	ChannelId      string `json:"channelId,omitempty"`
	Endpoint       string `json:"endpoint,omitempty"`
	Formats        string `json:"formats,omitempty"`
	OtherChannelId string `json:"otherChannelId,omitempty"`
	Originator     string `json:"originator,omitempty"`
	Timeout        int    `json:"timeout,omitempty"` // used in originate, not create
	/*
		These fields only used when connecting
		channel to a dialplan context.
		Use "App" to connect directly into ARI.
	*/

	Context   string `json:"context,omitempty"`
	Extension string `json:"extension,omitempty"`
	Label     string `json:"label,omitempty"`
	Priority  string `json:"priority,omitempty"`
}

type ChannelContinueOpts struct {
	Context   string `json:"context,omitempty"`
	Extension string `json:"extension,omitempty"`
	Priority  string `json:"priority,omitempty"`
	Label     string `json:"label,omitempty"`
}

type ChannelRecordOpts struct {
	MaxDurationSeconds int    `json:"maxDurationSeconds,omitempty"`
	MaxSilenceSeconds  int    `json:"maxSilenceSeconds,omitempty"`
	IfExists           string `json:"ifExists,omitempty"` // fail (default), overwrite, append
	Beep               bool   `json:"beep,omitempty"`
	TerminateOn        string `json:"terminateOn,omitempty"` // none (default), any, *, #

}

func (ch ChannelRecordOpts) Default() {
	ch.MaxDurationSeconds = 0
	ch.MaxSilenceSeconds = 0
	ch.IfExists = "fail"
	ch.Beep = false
	ch.TerminateOn = "none"
}

func (ch ChannelRecordOpts) Validate() error {
	allowedIfExists := []string{"fail", "overwrite", "append"}
	if !slices.Contains(allowedIfExists, ch.IfExists) {
		return fmt.Errorf("invalid value IfExists=%s, accepted: %s", ch.IfExists, strings.Join(allowedIfExists, ","))
	}

	allowedTerminate := []string{"none", "any", "*", "#"}
	if !slices.Contains(allowedTerminate, ch.TerminateOn) {
		return fmt.Errorf("invalid value TerminateOn=%s, accepted: %s", ch.TerminateOn, strings.Join(allowedTerminate, ","))
	}

	return nil
}
