package main

import (
	"context"
	"log"
	"time"

	"contrib.go.opencensus.io/exporter/zipkin"
	open_zipkin "github.com/openzipkin/zipkin-go"
	zipkin_http "github.com/openzipkin/zipkin-go/reporter/http"
	"go.opencensus.io/trace"
)

func main() {
	exporter, flush, err := NewExporter("sample", true)
	if err != nil {
		log.Fatal(err)
	}
	defer flush()
	trace.RegisterExporter(exporter)

	_, span := trace.StartSpan(context.Background(), "main")
	defer span.End()
	time.Sleep(1 * time.Second)
}

type FlushFunc func()

func NewExporter(service string, alwaysSample bool) (trace.Exporter, FlushFunc, error) {
	// 1. Configure exporter to export traces to Zipkin.
	localEndpoint, err := open_zipkin.NewEndpoint(service, "192.168.1.5:5454")
	if err != nil {
		return nil, nil, err
	}
	reporter := zipkin_http.NewReporter("http://localhost:9411/api/v2/spans")
	flush := func() { _ = reporter.Close() }
	exporter := zipkin.NewExporter(reporter, localEndpoint)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	return exporter, flush, nil
}
