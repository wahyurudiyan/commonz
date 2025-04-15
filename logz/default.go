package logz

import (
	"context"

	"go.uber.org/zap"
)

type fields []zap.Field

type defaultLogger struct {
	logger *zap.Logger
	fields []zap.Field
}

func (d *defaultLogger) setOtelMetadata(ctx context.Context) []zap.Field {
	var fields []zap.Field
	if hasSpanID(ctx) {
		fields = append(fields, zap.String("span-id", spanIdFromContext(ctx)))
	}

	if hasTraceID(ctx) {
		fields = append(fields, zap.String("span-id", traceIdFromContext(ctx)))
	}

	return fields
}

func (d *defaultLogger) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, d.setOtelMetadata(ctx)...)
	d.logger.Debug(msg, fields...)
}

func (d *defaultLogger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, d.setOtelMetadata(ctx)...)
	d.logger.Info(msg, fields...)
}

func (d *defaultLogger) Warning(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, d.setOtelMetadata(ctx)...)
	d.logger.Warn(msg, fields...)
}

func (d *defaultLogger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, d.setOtelMetadata(ctx)...)
	d.logger.Error(msg, fields...)
}

func (d *defaultLogger) DevelopmentPanic(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, d.setOtelMetadata(ctx)...)
	d.logger.DPanic(msg, fields...)
}

func (d *defaultLogger) Panic(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, d.setOtelMetadata(ctx)...)
	d.logger.Panic(msg, fields...)
}

func (d *defaultLogger) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, d.setOtelMetadata(ctx)...)
	d.logger.Fatal(msg, fields...)
}
