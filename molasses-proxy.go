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

func slowBan(r *http.Request, cst *goproxy.ProxyCtx) (*http.Request, *http.Response) {
  log.Println("banned request for", r.URL.String(), ", waiting:",request_number * 500,"milliseconds");
  time.Sleep(time.Duration(request_number) * time.Second/2)
  request_number += 1
  return r, nil
}

func main() {
  flag.Parse()
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

  log.Println("starting proxy on :", *port)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), proxy))
}
