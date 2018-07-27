package datadog

import (
	"context"
	"errors"

	godatadog "github.com/DataDog/opencensus-go-exporter-datadog"
	opencensus "github.com/devopsfaith/krakend-opencensus"
)

func init() {
	opencensus.RegisterExporterFactories(func(ctx context.Context, cfg opencensus.Config) (interface{}, error) {
		return Exporter(ctx, cfg)
	})
}

func Exporter(ctx context.Context, cfg opencensus.Config) (*godatadog.Exporter, error) {
	if cfg.Exporters.Datadog == nil {
		return nil, errors.New("datadog exporter disabled")
	}

	if cfg.Exporters.Datadog.Service == "" {
		cfg.Exporters.Datadog.Service = "KrakenD-opencensus"
	}

	e := godatadog.NewExporter(
		godatadog.Options{
			Namespace: cfg.Exporters.Datadog.Namespace,
			Service:   cfg.Exporters.Datadog.Service,
			TraceAddr: cfg.Exporters.Datadog.TraceAddr,
			StatsAddr: cfg.Exporters.Datadog.StatsAddr,
			Tags:      cfg.Exporters.Datadog.Tags,
		})

	go func() {
		<-ctx.Done()
		e.Stop()
	}()

	return e, nil
}
