package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptrace"

	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"

	"github.com/oinume/opencensus-client-trace-sample/opentelemetry"
)

func main() {
	tracerProvider, err := opentelemetry.NewTracerProvider("otelhttp_client_trace")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tracerProvider.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()
	otel.SetTracerProvider(tracerProvider)

	ctx := context.Background()
	ctx, span := tracerProvider.Tracer("main").Start(ctx, "main")
	defer span.End()

	if err := httpGet(ctx, "https://journal.lampetty.net/"); err != nil {
		log.Fatal(err)
	}
}

func httpGet(ctx context.Context, url string) error {
	ctx, span := otel.Tracer("main").Start(ctx, "httpGet")
	defer span.End()
	span.SetAttributes(attribute.Key("url").String(url))

	clientTrace := otelhttptrace.NewClientTrace(ctx)
	ctx = httptrace.WithClientTrace(ctx, clientTrace)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}
