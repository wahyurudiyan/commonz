package logz

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func setField(key string, val any) zap.Field {
	var f zap.Field
	switch val.(type) {
	case string:
		f = zap.String(key, val.(string))
	case int:
		f = zap.Int(key, val.(int))
	default:
		f = zap.Any(key, val)
	}

	return f
}

func MapToFields(fields map[string]any) []zap.Field {
	zf := func(items map[string]any) []zap.Field {
		var fields []zap.Field
		for k, v := range items {
			fields = append(fields, setField(k, v))
		}

		return fields
	}(fields)

	return zf
}

// Open Telemetry Basic Span and Tracer Data
func hasSpanID(ctx context.Context) bool {
	return trace.SpanContextFromContext(ctx).HasSpanID()
}

func hasTraceID(ctx context.Context) bool {
	return trace.SpanContextFromContext(ctx).HasTraceID()
}

func spanIdFromContext(ctx context.Context) string {
	sc := trace.SpanContextFromContext(ctx)
	return sc.SpanID().String()
}

func traceIdFromContext(ctx context.Context) string {
	sc := trace.SpanContextFromContext(ctx)
	return sc.TraceID().String()
}
