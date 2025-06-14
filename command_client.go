package ari

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
)

type CommandClient struct {
	host       string
	user       string
	pass       string
	httpClient *http.Client
	url        *url.URL
}

func NewCommandClient(host string, user string, pass string, timeout int) (*CommandClient, error) {
	t := time.Duration(timeout) * time.Second

	u, err := url.ParseRequestURI(host)
	if err != nil {
		return nil, err
	}

	return &CommandClient{
		host:       host,
		user:       user,
		pass:       pass,
		httpClient: &http.Client{Timeout: t},
		url:        u,
	}, nil

}

func (c *CommandClient) Url() string {
	return c.url.String()
}

func (c *CommandClient) httpGet(ctx context.Context, relUrl string) ([]byte, error) {
	return c.httpSend(ctx, http.MethodGet, relUrl, http.NoBody)

}

func (c *CommandClient) httpPost(ctx context.Context, relUrl string, body []byte) ([]byte, error) {

	var b io.Reader

	if body == nil {
		b = http.NoBody
	} else {
		b = bytes.NewReader(body)
	}

	return c.httpSend(ctx, http.MethodPost, relUrl, b)

}

func (c *CommandClient) httpDelete(ctx context.Context, relUrl string) ([]byte, error) {
	return c.httpSend(ctx, http.MethodDelete, relUrl, http.NoBody)

}
func (c *CommandClient) httpPut(ctx context.Context, relUrl string, body []byte) ([]byte, error) {
	var b io.Reader
	if body == nil {
		b = http.NoBody
	} else {
		b = bytes.NewReader(body)
	}
	return c.httpSend(ctx, http.MethodPut, relUrl, b)
}

func (c *CommandClient) httpSend(ctx context.Context, method string, relUrl string, body io.Reader) ([]byte, error) {
	// relUrl = path + query, technically called a "relative reference"
	// https://www.rfc-editor.org/rfc/rfc3986#section-4.2

	// parse the path into a url object so that we can extract the path + query
	p, err := url.Parse(relUrl)
	if err != nil {
		return nil, err
	}

	u := *c.url // make a shallow copy
	u.Path = path.Join(c.url.Path, p.Path)
	u.RawQuery = p.RawQuery

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.user, c.pass)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// check for HTTP error code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP %d %s: %s", resp.StatusCode, resp.Status, string(b))

	}

	b, err := io.ReadAll(resp.Body)

	return b, err

}
