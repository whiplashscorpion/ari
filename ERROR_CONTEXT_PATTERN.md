# Error Context Pattern

## Standard Patterns

### Path Building Errors
```go
path, err := url.JoinPath("/channels", channelId)
if err != nil {
    return fmt.Errorf("failed to build X path: %w", err)
}
```

### HTTP Operation Errors
```go
result, err := c.httpGet(ctx, path)
if err != nil {
    return fmt.Errorf("failed to X: %w", err)
}
```

### JSON Unmarshaling Errors
```go
err = json.Unmarshal(result, &output)
if err != nil {
    return fmt.Errorf("failed to unmarshal X: %w", err)
}
```

### Operation-Specific Examples

- **Channel operations**: "failed to answer channel", "failed to mute channel"
- **Bridge operations**: "failed to create bridge", "failed to add channel to bridge"
- **Asterisk operations**: "failed to get asterisk info", "failed to ping asterisk"
- **Application operations**: "failed to subscribe to application", "failed to get application"

## Required Import
```go
import "fmt"
```

## Implementation Notes
- Always use `%w` verb to wrap the original error
- Use descriptive context that indicates the specific operation
- Maintain consistent verb tense ("failed to...")
- Be specific about what failed (not just "operation failed")