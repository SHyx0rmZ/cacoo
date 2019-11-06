package cacoo

import (
	"context"
)

type apiKey int

var apiKeyKey apiKey

func WithApiKey(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, apiKeyKey, key)
}
