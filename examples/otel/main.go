package main

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"

	"github.com/oinume/opencensus-client-trace-sample/opentelemetry"
)

func main() {
	tracerProvider, err := opentelemetry.NewTracerProvider("otel")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tracerProvider.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()
	otel.SetTracerProvider(tracerProvider)

	ctx, span := tracerProvider.Tracer("main").Start(context.Background(), "main")
	defer span.End()
	func1(ctx)
	func2(ctx)
}

func func1(ctx context.Context) {
	_, span := otel.GetTracerProvider().Tracer("main").Start(ctx, "func1")
	defer span.End()
	time.Sleep(100 * time.Millisecond)
}

func func2(ctx context.Context) {
	_, span := otel.GetTracerProvider().Tracer("main").Start(ctx, "func2")
	defer span.End()
	time.Sleep(200 * time.Millisecond)
}
