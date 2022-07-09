module github.com/oinume/opencensus-client-trace-sample

go 1.16

require (
	contrib.go.opencensus.io/exporter/zipkin v0.1.2
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/openzipkin/zipkin-go v0.4.0
	go.opencensus.io v0.23.0
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.33.0
	go.opentelemetry.io/otel v1.8.0
	go.opentelemetry.io/otel/exporters/jaeger v1.8.0
	go.opentelemetry.io/otel/sdk v1.8.0
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	google.golang.org/grpc v1.47.0 // indirect
)
