package trace

import (
	"context"

	contextkeys "github.com/gladinov/contracts/context"
)

func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, contextkeys.TraceIDKey, traceID)
}

func TraceIDFromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(contextkeys.TraceIDKey).(string)
	return id, ok
}
