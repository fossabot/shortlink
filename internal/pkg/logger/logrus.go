package logger

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otellogrus"
	"go.opentelemetry.io/otel/attribute"

	"github.com/batazor/shortlink/internal/pkg/logger/field"
	"github.com/batazor/shortlink/internal/pkg/logger/tracer"
)

type logrusLogger struct {
	logger *logrus.Logger
}

func (log *logrusLogger) init(config Configuration) error {
	log.logger = logrus.New()

	// Logging =================================================================
	// Setup the logger backend using sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/sirupsen/logrus
	log.logger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: config.TimeFormat,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "msg",
			logrus.FieldKeyFunc:  "caller",
		},
	}

	// Tracing
	log.logger.AddHook(otellogrus.NewHook(otellogrus.WithLevels(
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	)))

	log.logger.SetReportCaller(false) // TODO: https://github.com/sirupsen/logrus/pull/973
	log.logger.SetOutput(config.Writer)
	log.setLogLevel(config.Level)

	return nil
}

func (log *logrusLogger) Close() error {
	return nil
}

func (log *logrusLogger) Get() interface{} {
	return log.logger
}

func (log *logrusLogger) converter(fields ...field.Fields) *logrus.Entry {
	logrusFields := logrus.Fields{}

	for _, field := range fields {
		for k, v := range field {
			logrusFields[k] = v
		}
	}

	entryLog := log.logger.WithFields(logrusFields)

	return entryLog
}

func (log *logrusLogger) setLogLevel(logLevel int) {
	switch logLevel {
	case FATAL_LEVEL:
		log.logger.SetLevel(logrus.FatalLevel)
	case ERROR_LEVEL:
		log.logger.SetLevel(logrus.ErrorLevel)
	case WARN_LEVEL:
		log.logger.SetLevel(logrus.WarnLevel)
	case INFO_LEVEL:
		log.logger.SetLevel(logrus.InfoLevel)
	case DEBUG_LEVEL:
		log.logger.SetLevel(logrus.DebugLevel)
	default:
		log.logger.SetLevel(logrus.InfoLevel)
	}
}

// Fatal ===============================================================================================================

func (log *logrusLogger) Fatal(msg string, fields ...field.Fields) {
	log.converter(fields...).Fatal(msg)
}

func (log *logrusLogger) FatalWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, nil, fields...)
	if err != nil {
		log.logger.WithContext(ctx).Error(fmt.Sprintf("Error send span to openTelemetry: %s", err.Error()))
	}

	log.converter(fields...).WithContext(ctx).Fatal(msg)
}

// Error ===============================================================================================================

func (log *logrusLogger) Error(msg string, fields ...field.Fields) {
	log.converter(fields...).Error(msg)
}

func (log *logrusLogger) ErrorWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	tags := []attribute.KeyValue{{
		Key:   "error",
		Value: attribute.BoolValue(true),
	}}

	fields, err := tracer.NewTraceFromContext(ctx, msg, tags, fields...)
	if err != nil {
		log.logger.WithContext(ctx).Error(fmt.Sprintf("Error send span to openTelemetry: %s", err.Error()))
	}

	log.converter(fields...).WithContext(ctx).Error(msg)
}

// Warn ================================================================================================================

func (log *logrusLogger) Warn(msg string, fields ...field.Fields) {
	log.converter(fields...).Warn(msg)
}

func (log *logrusLogger) WarnWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, nil, fields...)
	if err != nil {
		log.logger.WithContext(ctx).Error(fmt.Sprintf("Error send span to openTelemetry: %s", err.Error()))
	}

	log.converter(fields...).WithContext(ctx).Warn(msg)
}

// Info ================================================================================================================

func (log *logrusLogger) Info(msg string, fields ...field.Fields) {
	log.converter(fields...).Info(msg)
}

func (log *logrusLogger) InfoWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, nil, fields...)
	if err != nil {
		log.logger.WithContext(ctx).Error(fmt.Sprintf("Error send span to openTelemetry: %s", err.Error()))
	}

	log.converter(fields...).WithContext(ctx).Info(msg)
}

// Debug ===============================================================================================================

func (log *logrusLogger) Debug(msg string, fields ...field.Fields) {
	log.converter(fields...).Debug(msg)
}

func (log *logrusLogger) DebugWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, nil, fields...)
	if err != nil {
		log.logger.WithContext(ctx).Error(fmt.Sprintf("Error send span to openTelemetry: %s", err.Error()))
	}

	log.converter(fields...).WithContext(ctx).Debug(msg)
}
