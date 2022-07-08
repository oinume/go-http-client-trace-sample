# go-http-client-trace-sample

This is an example project for OpenCensus and httptrace.ClientTrace.
See [my blog post](https://journal.lampetty.net/entry/opencensus-httptrace) for details.

# Requirements

- Go
- Docker
- docker-compose

# How to run

```sh
$ docker-compose up
$ go run ./examples/trace/main.go
$ open http://localhost:16686/
```
