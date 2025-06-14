package ari

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

func (c *CommandClient) BridgesList(ctx context.Context) ([]Bridge, error) {
	var output []Bridge
	path := "/bridges"

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
func (c *CommandClient) BridgeCreate(ctx context.Context, bridgeId string, bridgeType string, name string) (Bridge, error) {
	path := "/bridges"
	params := url.Values{}

	if bridgeType != "" {
		params.Set("type", bridgeType)
	}

	if name != "" {
		params.Set("name", name)
	}

	path = path + "?" + params.Encode()

	result, err := c.httpPost(ctx, path, nil)
	if err != nil {
		return Bridge{}, err
	}

	var output Bridge
	err = json.Unmarshal(result, &output)
	return output, err

}
func (c *CommandClient) BridgeCreateWithId(ctx context.Context, bridgeId string, bridgeType string, name string) error {
	// documentation says that returns a Bridge object, but it does not.

	if bridgeId == "" {
		return fmt.Errorf("BridgeCreateWithId, empty param: bridgeId")
	}

	path, err := url.JoinPath("bridges", bridgeId)
	if err != nil {
		return err
	}

	params := url.Values{}

	if bridgeType != "" {
		params.Set("type", bridgeType)
	}

	if name != "" {
		params.Set("name", name)
	}

	path = path + "?" + params.Encode()

	_, err = c.httpPost(ctx, path, nil)

	return err

}

func (c *CommandClient) BridgeCreateUpdate(ctx context.Context, bridgeId string, opts ...BridgeCreateOpts) (Bridge, error) {
	var output Bridge
	path, err := url.JoinPath("/bridges", bridgeId)
	if err != nil {
		return output, err
	}

	if opts != nil {
		qparams := url.Values{}
		opts[0].formatQueryOpts(&qparams)
		path = path + "?" + qparams.Encode()
	}

	result, err := c.httpPost(ctx, path, nil)
	if err != nil {
		return output, err
	}

	err = json.Unmarshal(result, &output)
	if err != nil {
		return output, err
	}

	return output, nil

}
func (c *CommandClient) BridgeGet(ctx context.Context, bridgeId string) (Bridge, error) {
	var output Bridge
	path, err := url.JoinPath("/bridges", bridgeId)
	if err != nil {
		return output, err
	}

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
func (c *CommandClient) BridgesShutdown(ctx context.Context, bridgeId string) error {
	path, err := url.JoinPath("/bridges", bridgeId)
	if err != nil {
		return err
	}

	_, err = c.httpDelete(ctx, path)

	return err

}
func (c *CommandClient) BridgeAddChannel(ctx context.Context, bridgeId string, channel string, opts ...BridgeChannelOpts) error {
	path, err := url.JoinPath("/bridges", bridgeId, "addChannel")
	if err != nil {
		return err
	}

	qparams := url.Values{}
	qparams.Add("channel", channel)
	if opts != nil {
		opts[0].formatQueryOpts(&qparams)
		path = path + "?" + qparams.Encode()

	}

	_, err = c.httpPost(ctx, path, nil)

	return err

}
func (c *CommandClient) BridgesRemoveChannel(ctx context.Context, bridgeId string, channel string) error {
	path, err := url.JoinPath("/bridges", bridgeId, "removeChannel")
	if err != nil {
		return err
	}

	qparams := url.Values{}
	qparams.Add("channel", channel)
	path = path + "?" + qparams.Encode()

	_, err = c.httpPost(ctx, path, nil)
	return err

}

func (c *CommandClient) BridgesSetVideoSource(ctx context.Context, bridgeId string, channelId string) error {
	path, err := url.JoinPath("/bridges", bridgeId, "videoSource", channelId)
	if err != nil {
		return err
	}

	_, err = c.httpPost(ctx, path, nil)
	return err
}

func (c *CommandClient) BridgeRemoveVideoSrouce(ctx context.Context, bridgeId string) error {
	path, err := url.JoinPath("/bridges", bridgeId, "videoSource")
	if err != nil {
		return err
	}

	_, err = c.httpDelete(ctx, path)
	return err

}

func (c *CommandClient) BridgePlayMoh(ctx context.Context, bridgeId string, mohClass ...string) error {
	path, err := url.JoinPath("/bridges", bridgeId, "moh")
	if err != nil {
		return err
	}

	if mohClass != nil {
		qparams := url.Values{}
		qparams.Add("mohClass", mohClass[0])
		path = path + "?" + qparams.Encode()
	}

	_, err = c.httpPost(ctx, path, nil)
	return err

}

func (c *CommandClient) BridgeStopMoh(ctx context.Context, bridgeId string) error {
	path, err := url.JoinPath("/bridges", bridgeId, "moh")
	if err != nil {
		return err
	}

	_, err = c.httpDelete(ctx, path)
	return err

}

func (c *CommandClient) BridgePlay(ctx context.Context, bridgeId string, media string, opts ...BridgePlaybackOpts) (Playback, error) {
	var output Playback

	path, err := url.JoinPath("/bridges", bridgeId, "play")
	if err != nil {
		return output, err
	}

	qparams := url.Values{}
	qparams.Add("media", media)
	if opts != nil {
		opts[0].formatQueryOpts(&qparams)
	}
	path = path + "?" + qparams.Encode()

	result, err := c.httpPost(ctx, path, nil)
	if err != nil {
		return output, err
	}

	err = json.Unmarshal(result, &output)
	if err != nil {
		return output, err
	}
	return output, nil

}

func (c *CommandClient) BridgePlayWithId(ctx context.Context, bridgeId string, media string, playbackId string, opts ...BridgePlaybackOpts) (Playback, error) {
	var output Playback

	path, err := url.JoinPath("/bridges", bridgeId, "play", playbackId)
	if err != nil {
		return output, err
	}

	qparams := url.Values{}
	qparams.Add("media", media)
	if opts != nil {
		opts[0].formatQueryOpts(&qparams)
	}

	qparams.Del("playbackId") // override the opts, playbackId is a path param in this function
	path = path + "?" + qparams.Encode()

	result, err := c.httpPost(ctx, path, nil)
	if err != nil {
		return output, err
	}

	err = json.Unmarshal(result, &output)
	return output, err

}

func (c *CommandClient) BridgeRecord(ctx context.Context, bridgeId string, name string, format string, opts ...BridgeRecordOpts) (LiveRecording, error) {
	var output LiveRecording

	path, err := url.JoinPath("/bridges", bridgeId, "record")
	if err != nil {
		return output, err
	}

	qparams := url.Values{}
	qparams.Add("name", name)
	qparams.Add("format", format)
	if opts != nil {
		opts[0].formatQueryOpts(&qparams)
	}
	path = path + "?" + qparams.Encode()

	result, err := c.httpPost(ctx, path, nil)
	if err != nil {
		return output, err
	}

	err = json.Unmarshal(result, &output)
	return output, err

}

type BridgeCreateOpts struct {
	BridgeId string `json:"bridgeId,omitempty"`
	Type     string `json:"type,omitempty"`
	Name     string `json:"name,omitempty"`
}

func (opts BridgeCreateOpts) formatQueryOpts(qparams *url.Values) {
	params := make(map[string]string)

	if opts.BridgeId != "" {
		params["bridgeId"] = opts.BridgeId
	}
	if opts.Type != "" {
		params["type"] = opts.Type
	}
	if opts.Type != "" {
		params["name"] = opts.Name
	}

	if len(params) != 0 {
		for key, val := range params {
			qparams.Add(key, val)
		}
	}
}

type BridgeChannelOpts struct {
	// *bool used to allow for omitting optional fields
	Role                        string `json:"role,omitempty"`
	AbsorbDTMF                  *bool  `json:"absorbDTMF,omitempty"`
	Mute                        *bool  `json:"mute,omitempty"`
	InhibitConnectedLineUpdates *bool  `json:"inhibitConnectedLineUpdates,omitempty"`
}

func (opts BridgeChannelOpts) formatQueryOpts(qparams *url.Values) {
	params := make(map[string]string)

	if opts.Role != "" {
		params["role"] = opts.Role
	}
	if opts.AbsorbDTMF != nil {
		params["absorbDTMF"] = strconv.FormatBool(*opts.AbsorbDTMF)
	}
	if opts.Mute != nil {
		params["mute"] = strconv.FormatBool(*opts.AbsorbDTMF)

	}
	if opts.InhibitConnectedLineUpdates != nil {
		params["inhibitConnectedLineUpdates"] = strconv.FormatBool(*opts.InhibitConnectedLineUpdates)

	}
	for key, val := range params {
		qparams.Add(key, val)
	}
}

type BridgePlaybackOpts struct {
	PlaybackId string
	Lang       string
	Offsetms   int
	Skipms     int
}

func (opts BridgePlaybackOpts) formatQueryOpts(qparams *url.Values) {
	params := make(map[string]string)

	if opts.PlaybackId != "" {
		params["playbackId"] = opts.PlaybackId
	}
	if opts.Lang != "" {
		params["lang"] = opts.Lang
	}
	if opts.Offsetms != 0 {
		params["offsetms"] = strconv.Itoa(opts.Offsetms)
	}
	if opts.Skipms != 0 {
		params["skipms"] = strconv.Itoa(opts.Skipms)
	}
	for key, val := range params {
		qparams.Add(key, val)
	}

}

type BridgeRecordOpts struct {
	MaxDurationSeconds int
	MaxSilenceSeconds  int
	IfExists           string
	Beep               *bool
	TerminateOn        string
}

func (opts BridgeRecordOpts) formatQueryOpts(qparams *url.Values) {
	params := make(map[string]string)

	if opts.MaxDurationSeconds != 0 {
		params["maxDurationSeconds"] = strconv.Itoa(opts.MaxDurationSeconds)
	}
	if opts.MaxSilenceSeconds != 0 {
		params["maxSilenceSeconds"] = strconv.Itoa(opts.MaxSilenceSeconds)
	}
	if opts.IfExists != "" {
		params["ifExists"] = opts.IfExists
	}
	if opts.Beep != nil {
		params["beep"] = strconv.FormatBool(*opts.Beep)
	}
	if opts.TerminateOn != "" {
		params["terminateOn"] = opts.TerminateOn
	}

	for key, val := range params {
		qparams.Add(key, val)
	}
}
