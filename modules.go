package ari

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

func (c *CommandClient) ModulesList(ctx context.Context) ([]Module, error) {
	var output []Module
	path := "/asterisk/modules"

	result, err := c.httpGet(ctx, path)
	if err != nil {
		return output, fmt.Errorf("failed to get modules: %w", err)
	}

	err = json.Unmarshal(result, &output)
	if err != nil {
		return output, fmt.Errorf("failed to unmarshal modules: %w", err)
	}

	return output, nil

}

func (c *CommandClient) ModulesGet(ctx context.Context, moduleName string) (Module, error) {
	var output Module

	path, err := url.JoinPath("/asterisk/modules", moduleName)
	if err != nil {
		return output, fmt.Errorf("failed to build module path: %w", err)
	}

	result, err := c.httpGet(ctx, path)
	if err != nil {
		return output, fmt.Errorf("failed to get module: %w", err)

	}

	err = json.Unmarshal(result, &output)
	if err != nil {
		return output, fmt.Errorf("failed to unmarshal module: %w", err)
	}

	return output, nil

}

func (c *CommandClient) ModulesLoad(ctx context.Context, moduleName string) error {
	path, err := url.JoinPath("/asterisk/modules", moduleName)
	if err != nil {
		return fmt.Errorf("failed to build module path: %w", err)
	}

	_, err = c.httpPost(ctx, path, nil)
	if err != nil {
		return fmt.Errorf("failed to load module: %w", err)
	}

	return nil
}
func (c *CommandClient) ModulesUnload(ctx context.Context, moduleName string) error {
	path, err := url.JoinPath("/asterisk/modules", moduleName)
	if err != nil {
		return fmt.Errorf("failed to build module path: %w", err)
	}

	_, err = c.httpDelete(ctx, path)
	if err != nil {
		return fmt.Errorf("failed to unload module: %w", err)
	}
	return nil
}

func (c *CommandClient) ModulesReload(ctx context.Context, moduleName string) error {
	path, err := url.JoinPath("/asterisk/modules", moduleName)
	if err != nil {
		return fmt.Errorf("failed to build module path: %w", err)
	}

	_, err = c.httpPut(ctx, path, nil)
	if err != nil {
		return fmt.Errorf("failed to reload module: %w", err)
	}

	return nil
}
