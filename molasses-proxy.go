package main

import (
	"flag"
	"fmt"
	"github.com/elazarl/goproxy"
	"github.com/ox/molasses-proxy/linkio"
	"log"
	"net/http"
)

var port = flag.Int("port", 8080,
	"the port to listen for requests on")
var help = flag.Bool("help", false,
	"print this help message")
var rate = flag.Int("rate", 56,
	"set the maximum link rate in kbps")

func main() {
	flag.Parse()
	slowLink := linkio.NewLink(*rate)

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	proxy.OnResponse().DoFunc(
		func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			resp.Body = slowLink.NewLinkReader(resp.Body)
			return resp
		})

	if *help {
		fmt.Println("usage:\n\n\tmolasses-proxy [--port=8080] [--rate=0]\n")
		flag.PrintDefaults()
		return
	}

	log.Println("starting proxy on :", *port, "@", *rate, "kbps")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), proxy))
}
