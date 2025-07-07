package otel

import (
	"context"

	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

type Config struct {
	ServerName    string `yaml:"name" json:"name,omitempty"`
	ServerVersion string `yaml:"version" json:"version,omitempty"`
	Endpoint      string `yaml:"endpoint" json:"endpoint,omitempty"`
}

type Provider struct {
	cfg Config
	mp  *metric.MeterProvider
	tp  *trace.TracerProvider
}

func NewProvider(ctx context.Context, cfg Config) (*Provider, error) {
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(cfg.ServerName),
		semconv.ServiceVersion(cfg.ServerVersion),
	)

	mp, err := newMeterProvider(ctx, cfg.Endpoint, res)
	if err != nil {
		return nil, err
	}

	tp, err := newTracerProvider(ctx, cfg.Endpoint, res)
	if err != nil {
		return nil, err
	}

	return &Provider{cfg: cfg, mp: mp, tp: tp}, nil
}

func (p *Provider) Shutdown(ctx context.Context) error {
	if p.mp != nil {
		if err := p.mp.Shutdown(ctx); err != nil {
			return err
		}
	}

	if p.tp != nil {
		if err := p.tp.Shutdown(ctx); err != nil {
			return err
		}
	}

	return nil
}
