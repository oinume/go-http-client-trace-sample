package opentelemetry

import (
	_ "go.opencensus.io/resource"
	_ "go.opencensus.io/trace"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

func NewTracerProvider(serviceName string) (*trace.TracerProvider, error) {
	// Port details: https://www.jaegertracing.io/docs/getting-started/
	collectorEndpointURI := "http://localhost:14268/api/traces"

	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(collectorEndpointURI)))
	if err != nil {
		return nil, err
	}

	r := NewResource(serviceName, "v1", "local")
	return trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(r),
		trace.WithSampler(trace.TraceIDRatioBased(1)),
	), nil
}

func NewResource(serviceName string, version string, environment string) *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
			semconv.ServiceVersionKey.String(version),
			attribute.String("environment", environment),
		),
	)
	return r
}
