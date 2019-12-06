package main

import (
	"context"
	"log"
	"time"

	"github.com/oinume/opencensus-client-trace-sample/opencensus"
	"go.opencensus.io/trace"
)

func main() {
	exporter, flush, err := opencensus.NewExporter("sample")
	if err != nil {
		log.Fatal(err)
	}
	defer flush()
	trace.RegisterExporter(exporter)

	ctx, span := trace.StartSpan(context.Background(), "main")
	defer span.End()
	func1(ctx)
	func2(ctx)
}

func func1(ctx context.Context) {
	_, span := trace.StartSpan(ctx, "func1")
	defer span.End()
	time.Sleep(100 * time.Millisecond)
}

func func2(ctx context.Context) {
	_, span := trace.StartSpan(ctx, "func2")
	defer span.End()
	time.Sleep(200 * time.Millisecond)
}
