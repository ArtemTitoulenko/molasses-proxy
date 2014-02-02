package main

import (
	"flag"
	"fmt"
	"github.com/elazarl/goproxy"
	"github.com/ox/molasses-proxy/linkio"
	"io/ioutil"
	"log"
	"net/http"
)

var port = flag.Int("port", 8080,
	"the port to listen for requests on")
var help = flag.Bool("help", false,
	"print this help message")
var rate = flag.Int("rate", 56,
	"the maximum link rate in kbps")

func main() {
	flag.Parse()
	slowLink := linkio.NewLink(*rate)

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	proxy.OnRequest().DoFunc(
		func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			req.Body = ioutil.NopCloser(slowLink.NewLinkReader(req.Body))
			return req, nil
		})

	proxy.OnResponse().DoFunc(
		func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			if resp.Body != nil {
				resp.Body = ioutil.NopCloser(slowLink.NewLinkReader(resp.Body))
			}

			return resp
		})

	if *help {
		fmt.Println("usage:\n\n\tmolasses-proxy [--port=8080] [--rate=56]\n")
		flag.PrintDefaults()
		return
	}

	log.Println("starting proxy on :", *port, "@", *rate, "kbps")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), proxy))
}
