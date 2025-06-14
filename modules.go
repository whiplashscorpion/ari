package ari

import (
	"context"
	"encoding/json"
	"net/url"
)

func (c *CommandClient) ModulesList(ctx context.Context) ([]Module, error) {
	var output []Module
	path := "/asterisk/modules"

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

func (c *CommandClient) ModulesGet(ctx context.Context, moduleName string) (Module, error) {
	var output Module

	path, err := url.JoinPath("/asterisk/modules", moduleName)
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

func (c *CommandClient) ModulesLoad(ctx context.Context, moduleName string) error {
	path, err := url.JoinPath("/asterisk/modules", moduleName)
	if err != nil {
		return err
	}

	_, err = c.httpPost(ctx, path, nil)
	if err != nil {
		return err
	}

	return nil
}
func (c *CommandClient) ModulesUnload(ctx context.Context, moduleName string) error {
	path, err := url.JoinPath("/asterisk/modules", moduleName)
	if err != nil {
		return err
	}

	_, err = c.httpDelete(ctx, path)
	return err
}

func (c *CommandClient) ModulesReload(ctx context.Context, moduleName string) error {
	path, err := url.JoinPath("/asterisk/modules", moduleName)
	if err != nil {
		return err
	}

	_, err = c.httpPut(ctx, path, nil)
	if err != nil {
		return err
	}

	return nil
}
