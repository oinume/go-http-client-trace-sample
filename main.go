package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptrace"

	"github.com/oinume/opencensus-client-trace-sample/opencensus"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

func main() {
	// TODO: Stackdriver exporter
	exporter, flush, err := opencensus.NewExporter("sample")
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
