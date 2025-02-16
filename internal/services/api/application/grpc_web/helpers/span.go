package helpers

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// RegisterSpan - inject spanId in response header
func RegisterSpan(ctx context.Context) error {
	return grpc.SendHeader(ctx, metadata.New(map[string]string{
		"trace-id": trace.LinkFromContext(ctx).SpanContext.TraceID().String(),
	}))
}
