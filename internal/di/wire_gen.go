// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"context"
	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/mq"
	"github.com/batazor/shortlink/internal/mq/kafka"
	"github.com/batazor/shortlink/internal/store"
	"github.com/batazor/shortlink/internal/traicing"
	"github.com/google/wire"
	"github.com/heptiolabs/healthcheck"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

// Injectors from wire.go:

func InitializeFullService(ctx context.Context) (*Service, func(), error) {
	logger, cleanup, err := InitLogger(ctx)
	if err != nil {
		return nil, nil, err
	}
	mq, err := InitMQ(ctx)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	serveMux := InitMonitoring()
	tracer, cleanup2, err := InitTracer(ctx, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	db, cleanup3, err := InitStore(ctx, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := NewFullService(logger, mq, serveMux, tracer, db)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

func InitializeLoggerService(ctx context.Context) (*Service, func(), error) {
	logger, cleanup, err := InitLogger(ctx)
	if err != nil {
		return nil, nil, err
	}
	mq, err := InitMQ(ctx)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	service, err := NewLoggerService(logger, mq)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup()
	}, nil
}

// wire.go:

// Service - heplers
type Service struct {
	Log    logger.Logger
	Tracer opentracing.Tracer
	// TracerClose func()
	DB         store.DB
	MQ         mq.MQ
	Monitoring *http.ServeMux
}

// InitStore return store
func InitStore(ctx context.Context, log logger.Logger) (store.DB, func(), error) {
	var st store.Store
	db := st.Use(ctx, log)

	cleanup := func() {
		if err := db.Close(); err != nil {
			log.Error(err.Error())
		}
	}

	return db, cleanup, nil
}

func InitLogger(ctx context.Context) (logger.Logger, func(), error) {
	viper.SetDefault("LOG_LEVEL", logger.INFO_LEVEL)
	viper.SetDefault("LOG_TIME_FORMAT", time.RFC3339Nano)

	conf := logger.Configuration{
		Level:      viper.GetInt("LOG_LEVEL"),
		TimeFormat: viper.GetString("LOG_TIME_FORMAT"),
	}

	log, err := logger.NewLogger(logger.Zap, conf)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {

		log.Close()
	}

	return log, cleanup, nil
}

func InitTracer(ctx context.Context, log logger.Logger) (opentracing.Tracer, func(), error) {
	viper.SetDefault("TRACER_SERVICE_NAME", "ShortLink")
	viper.SetDefault("TRACER_URI", "localhost:6831")

	config := traicing.Config{
		ServiceName: viper.GetString("TRACER_SERVICE_NAME"),
		URI:         viper.GetString("TRACER_URI"),
	}

	tracer, tracerClose, err := traicing.Init(config)
	if err != nil {
		return nil, nil, err
	}

	ctx = traicing.WithTraicer(ctx, tracer)

	cleanup := func() {
		if err := tracerClose.Close(); err != nil {
			log.Error(err.Error())
		}
	}

	return tracer, cleanup, nil
}

func InitMonitoring() *http.ServeMux {

	registry := prometheus.NewRegistry()

	health := healthcheck.NewMetricsHandler(registry, "common")

	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))

	commonMux := http.NewServeMux()

	commonMux.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))

	commonMux.HandleFunc("/live", health.LiveEndpoint)

	commonMux.HandleFunc("/ready", health.ReadyEndpoint)

	return commonMux
}

func InitMQ(ctx context.Context) (mq.MQ, error) {
	viper.SetDefault("MQ_ENABLED", "false")

	if viper.GetBool("MQ_ENABLED") {
		var service mq.MQ
		service = &kafka.Kafka{}

		if err := service.Init(ctx); err != nil {
			return nil, err
		}

		return service, nil
	}

	return nil, nil
}

// Default =============================================================================================================
var DefaultSet = wire.NewSet(InitLogger, InitTracer)

// FullService =========================================================================================================
var FullSet = wire.NewSet(DefaultSet, NewFullService, InitStore, InitMonitoring, InitMQ)

func NewFullService(log logger.Logger, mq2 mq.MQ, monitoring *http.ServeMux, tracer opentracing.Tracer, db store.DB) (*Service, error) {
	return &Service{
		Log:    log,
		MQ:     mq2,
		Tracer: tracer,

		Monitoring: monitoring,
		DB:         db,
	}, nil
}

// LoggerService =======================================================================================================
var LoggerSet = wire.NewSet(DefaultSet, NewLoggerService, InitMQ)

func NewLoggerService(log logger.Logger, mq2 mq.MQ) (*Service, error) {
	return &Service{
		Log: log,
		MQ:  mq2,
	}, nil
}
