package logger

import (
	"context"
	"fmt"
	"time"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/batazor/shortlink/internal/pkg/logger/field"
	"github.com/batazor/shortlink/internal/pkg/logger/tracer"
)

type zapLogger struct {
	logger *otelzap.Logger
}

func (log *zapLogger) init(config Configuration) error {
	logLevel := log.setLogLevel(config.Level)

	// To keep the example deterministic, disable timestamps in the output.
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoder(log.timeEncoder(config.TimeFormat)),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Wrap zap logger to extend Zap with API that accepts a context.Context.
	log.logger = otelzap.New(zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(zapcore.AddSync(config.Writer)),
		logLevel,
	), zap.AddCaller(), zap.AddCallerSkip(1)))

	return nil
}

func (log *zapLogger) Close() error {
	err := log.logger.Sync()
	return err
}

func (log *zapLogger) Get() interface{} {
	return log.logger
}

func (log *zapLogger) converter(fields ...field.Fields) []zap.Field {
	var zapFields []zap.Field

	for _, field := range fields {
		for k, v := range field {
			zapFields = append(zapFields, zap.Any(k, v))
		}
	}

	return zapFields
}

func (log *zapLogger) setLogLevel(logLevel int) zap.AtomicLevel {
	atom := zap.NewAtomicLevel()

	switch logLevel {
	case FATAL_LEVEL:
		atom.SetLevel(zapcore.Level(zap.FatalLevel))
	case ERROR_LEVEL:
		atom.SetLevel(zapcore.ErrorLevel)
	case WARN_LEVEL:
		atom.SetLevel(zapcore.WarnLevel)
	case INFO_LEVEL:
		atom.SetLevel(zapcore.InfoLevel)
	case DEBUG_LEVEL:
		atom.SetLevel(zapcore.DebugLevel)
	default:
		atom.SetLevel(zapcore.InfoLevel)
	}

	return atom
}

func (log *zapLogger) timeEncoder(format string) func(time.Time, zapcore.PrimitiveArrayEncoder) {
	return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(format))
	}
}

// Fatal ===============================================================================================================

func (log *zapLogger) Fatal(msg string, fields ...field.Fields) {
	zapFields := log.converter(fields...)
	log.logger.Fatal(msg, zapFields...)
}

func (log *zapLogger) FatalWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, nil, fields...)
	if err != nil {
		log.logger.Ctx(ctx).Error(fmt.Sprintf("Error send span to openTelemetry: %s", err.Error()))
	}

	zapFields := log.converter(fields...)
	log.logger.Ctx(ctx).Fatal(msg, zapFields...)
}

// Warn ================================================================================================================

func (log *zapLogger) Warn(msg string, fields ...field.Fields) {
	zapFields := log.converter(fields...)
	log.logger.Warn(msg, zapFields...)
}

func (log *zapLogger) WarnWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, nil, fields...)
	if err != nil {
		log.logger.Ctx(ctx).Error(fmt.Sprintf("Error send span to openTelemetry: %s", err.Error()))
	}

	zapFields := log.converter(fields...)
	log.logger.Ctx(ctx).Warn(msg, zapFields...)
}

// Error ===============================================================================================================

func (log *zapLogger) Error(msg string, fields ...field.Fields) {
	zapFields := log.converter(fields...)
	log.logger.Error(msg, zapFields...)
}

func (log *zapLogger) ErrorWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	tags := []attribute.KeyValue{{
		Key:   "error",
		Value: attribute.BoolValue(true),
	}}

	fields, err := tracer.NewTraceFromContext(ctx, msg, tags, fields...)
	if err != nil {
		log.logger.Ctx(ctx).Error(fmt.Sprintf("Error send span to openTelemetry: %s", err.Error()))
	}

	zapFields := log.converter(fields...)
	log.logger.Ctx(ctx).Error(msg, zapFields...)
}

// Info ================================================================================================================

func (log *zapLogger) Info(msg string, fields ...field.Fields) {
	zapFields := log.converter(fields...)
	log.logger.Info(msg, zapFields...)
}

func (log *zapLogger) InfoWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, nil, fields...)
	if err != nil {
		log.logger.Ctx(ctx).Error(fmt.Sprintf("Error send span to openTelemetry: %s", err.Error()))
	}

	zapFields := log.converter(fields...)
	log.logger.Ctx(ctx).Info(msg, zapFields...)
}

// Debug ===============================================================================================================

func (log *zapLogger) Debug(msg string, fields ...field.Fields) {
	zapFields := log.converter(fields...)
	log.logger.Debug(msg, zapFields...)
}

func (log *zapLogger) DebugWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, nil, fields...)
	if err != nil {
		log.logger.Ctx(ctx).Error(fmt.Sprintf("Error send span to openTelemetry: %s", err.Error()))
	}

	zapFields := log.converter(fields...)
	log.logger.Ctx(ctx).Debug(msg, zapFields...)
}
