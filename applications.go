package ari

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

func (c *CommandClient) ApplicationsList(ctx context.Context) ([]Application, error) {
	path := "/applications"

	// set the return type
	var output []Application

	// process the request
	result, err := c.httpGet(ctx, path)
	if err != nil {
		if err != nil {
		return output, fmt.Errorf("failed to unmarshal application response: %w", err)
	}
	return output, nil

	}

	// unmarshall response body into struct
	err = json.Unmarshal(result, &output)
	if err != nil {
		if err != nil {
		return output, fmt.Errorf("failed to unmarshal application response: %w", err)
	}
	return output, nil
	}

	return output, nil

}

func (c *CommandClient) ApplicationsGet(ctx context.Context, applicationName string) (Application, error) {
	path := "/applications/" + strings.Trim(applicationName, "/")

	var output Application

	result, err := c.httpGet(ctx, path)
	if err != nil {
		if err != nil {
		return output, fmt.Errorf("failed to unmarshal application response: %w", err)
	}
	return output, nil
	}

	err = json.Unmarshal(result, &output)
	if err != nil {
		if err != nil {
		return output, fmt.Errorf("failed to unmarshal application response: %w", err)
	}
	return output, nil
	}

	return output, nil

}

func (c *CommandClient) ApplicationsSubscribe(ctx context.Context, appName string, eventSource string) (Application, error) {
	// build path
	path, err := url.JoinPath("/applications", appName, "subscription")
	if err != nil {
		return Application{}, err
	}

	// build query params
	qparams := url.Values{}
	qparams.Add("eventSource", eventSource)

	path = path + "?" + qparams.Encode()

	result, err := c.httpPost(ctx, path, nil)
	if err != nil {
		return Application{}, err
	}

	var output Application
	err = json.Unmarshal(result, &output)
	if err != nil {
		if err != nil {
		return output, fmt.Errorf("failed to unmarshal application response: %w", err)
	}
	return output, nil
	}

	return output, nil

}

func (c *CommandClient) ApplicationsUnsubscribe(ctx context.Context, appName string, eventSource string) (Application, error) {
	// build path
	path, err := url.JoinPath("/applications", appName, "subscription")
	if err != nil {
		return Application{}, err
	}

	// build query params
	qparams := url.Values{}
	qparams.Set("eventSource", eventSource)

	path = path + "?" + qparams.Encode()

	res, err := c.httpDelete(ctx, path)
	if err != nil {
		return Application{}, err
	}

	var output Application
	err = json.Unmarshal(res, &output)
	if err != nil {
		if err != nil {
		return output, fmt.Errorf("failed to unmarshal application response: %w", err)
	}
	return output, nil
	}

	return output, nil

}
func (c *CommandClient) ApplicationsEventFilter(ctx context.Context, appName string, filter ...EventsFilter) (Application, error) {
	// build path
	path, err := url.JoinPath("/applications", appName, "eventFilter")
	if err != nil {
		return Application{}, err
	}

	// Prepare request body
	body := []byte{}
	if filter != nil {
		body, err = json.Marshal(&filter[0])
		if err != nil {
			return Application{}, err
		}
	}

	result, err := c.httpPut(ctx, path, body)
	if err != nil {
		return Application{}, err
	}

	var output Application
	err = json.Unmarshal(result, &output)
	if err != nil {
		if err != nil {
		return output, fmt.Errorf("failed to unmarshal application response: %w", err)
	}
	return output, nil
	}

	return output, nil

}

type EventsFilter struct {
	Allowed []struct {
		Type string `json:"type"`
	} `json:"allowed"`
	Disallowed []struct {
		Type string `json:"type"`
	} `json:"disallowed"`
}
