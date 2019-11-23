package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptrace"

	"contrib.go.opencensus.io/exporter/zipkin"
	open_zipkin "github.com/openzipkin/zipkin-go"
	zipkin_http "github.com/openzipkin/zipkin-go/reporter/http"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

func main() {
	// TODO: Stackdriver exporter
	exporter, flush, err := NewExporter("sample")
	if err != nil {
		log.Fatal(err)
	}
	defer flush()
	trace.RegisterExporter(exporter)

	ctx, span := trace.StartSpan(context.Background(), "main")
	defer span.End()

	if err := httpGet(ctx, "https://journal.lampetty.net/"); err != nil {
		log.Fatal(err)
	}
	//time.Sleep(1 * time.Second)
}

type FlushFunc func()

func NewExporter(service string) (trace.Exporter, FlushFunc, error) {
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

func httpGet(ctx context.Context, url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	ctx, span := trace.StartSpan(ctx, "httpGet")
	defer span.End()
	clientTrace := ochttp.NewSpanAnnotatingClientTrace(req, span)
	ctx = httptrace.WithClientTrace(ctx, clientTrace)
	req = req.WithContext(ctx)
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}
