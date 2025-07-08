package otel

import (
	"context"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

type Provider struct {
	opts   *Options
	meter  metric.MeterProvider
	tracer trace.TracerProvider
}

func NewProvider(ctx context.Context, opts ...Option) (*Provider, error) {
	opt := newOptions()
	for _, v := range opts {
		v(opt)
	}

	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(opt.serviceName),
		semconv.ServiceVersion(opt.serviceVersion),
	)

	mp, err := newMeterProvider(ctx, res, opt)
	if err != nil {
		return nil, err
	}

	tp, err := newTracerProvider(ctx, res, opt)
	if err != nil {
		return nil, err
	}

	return &Provider{
		opts:   opt,
		meter:  mp,
		tracer: tp,
	}, nil
}
