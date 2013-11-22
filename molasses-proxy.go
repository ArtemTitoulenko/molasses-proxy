package main

import (
  "github.com/elazarl/goproxy"
  "log"
  "net/http"
  "time"
  "os"
  "io/ioutil"
  "strings"
  "flag"
  "fmt"
)

var blocked_filename = "./blocked_hosts"
var request_number = 0
var port = flag.Int("port", 8080, "the port to listen for requests on")
var delay = flag.Int("delayms", 500, "increase the delay by this many milliseconds per request")
var help = flag.Bool("help", false, "print this help message")

func slowBan(r *http.Request, cst *goproxy.ProxyCtx) (*http.Request, *http.Response) {
  log.Println("banned request for", r.URL.String(), ", waiting:",request_number * (*delay),"milliseconds");
  time.Sleep(time.Duration(request_number * (*delay)))
  request_number += 1
  return r, nil
}

func main() {
  flag.Parse()

  if *help {
    fmt.Println("usage:\n\n\tmolasses-proxy [--port=8080] [--delayms=500]\n")
    flag.PrintDefaults()
    return
  }


  proxy := goproxy.NewProxyHttpServer()
  proxy.Verbose = false

  if _, err := os.Stat(blocked_filename); os.IsNotExist(err) {
    log.Fatalf("no such file or directory: %s\n", blocked_filename)
  }

  data, err := ioutil.ReadFile(blocked_filename)
  if err != nil {
    log.Fatalf("error reading %s: %s\n", blocked_filename, err)
  }

  websites := strings.Split(string(data), "\n")
  for _, website := range websites {
    if website == "" {
      break
    }

    log.Println("slow banning", website)
    proxy.OnRequest(goproxy.DstHostIs(website)).DoFunc(slowBan)
  }

  log.Println("starting proxy on :", *port, "delaying", *delay, "more milliseconds per request")
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), proxy))
}
