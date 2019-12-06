package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"net/textproto"
)

func main() {
	clientTrace := &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			fmt.Println("GetConn")
		},
		GotConn: func(info httptrace.GotConnInfo) {
			fmt.Println("GotConn")
		},
		PutIdleConn: func(err error) {
			fmt.Println("PutIdleConn")
		},
		GotFirstResponseByte: func() {
			fmt.Println("GotFirstResponseByte")
		},
		Got100Continue: func() {
			fmt.Println("Got100Continue")
		},
		Got1xxResponse: func(code int, header textproto.MIMEHeader) error {
			fmt.Println("Got1xxResponse")
			return nil
		},
		DNSStart: func(_ httptrace.DNSStartInfo) {
			fmt.Println("DNSStart")
		},
		DNSDone: func(_ httptrace.DNSDoneInfo) {
			fmt.Println("DNSDone")
		},
		ConnectStart: func(network, addr string) {
			fmt.Println("ConnectStart")
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Println("ConnectDone")
		},
		TLSHandshakeStart: func() {
			fmt.Println("TLSHandshakeStart")
		},
		TLSHandshakeDone: func(tls.ConnectionState, error) {
			fmt.Println("TLSHandshakeDone")
		},
		WroteHeaderField: func(key string, value []string) {
			fmt.Printf("WroteHeaderField: %v\n", key)
		},
		WroteHeaders: func() {
			fmt.Println("WroteHeaders")
		},
		Wait100Continue: func() {
			fmt.Println("Wait100Continue")
		},
		WroteRequest: func(httptrace.WroteRequestInfo) {
			fmt.Println("WroteRequest")
		},
	}

	req, err := http.NewRequest("GET", "https://journal.lampetty.net/", nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx := httptrace.WithClientTrace(context.Background(), clientTrace)
	req = req.WithContext(ctx)
	if _, err := http.DefaultClient.Do(req); err != nil {
		log.Fatal(err)
	}
}
