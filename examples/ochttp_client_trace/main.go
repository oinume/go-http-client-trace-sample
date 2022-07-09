package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptrace"

	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"

	"github.com/oinume/opencensus-client-trace-sample/opencensus"
)

func main() {
	exporter, flush, err := opencensus.NewExporter("ochttp_client_trace")
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
}

func httpGet(ctx context.Context, url string) error {
	ctx, span := trace.StartSpan(ctx, "httpGet")
	defer span.End()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	clientTrace := ochttp.NewSpanAnnotatingClientTrace(req, span)
	ctx = httptrace.WithClientTrace(ctx, clientTrace)
	req = req.WithContext(ctx)
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}
