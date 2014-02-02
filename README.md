# Molasses Proxy

Molasses Proxy is a bandwidth-limiting proxy. It can be used to constrain the bandwidth of all HTTP requests passing though it.

## Installation

    $ go get github.com/ArtemTitoulenko/molasses-proxy


## Running

    $ molasses-proxy

## Getting Help

    $ molasses-proxy --help
      usage:

          molasses-proxy [--port=8080] [--rate=56]

	  -help=false: print this help message
	  -port=8080: the port to listen for requests on
	  -rate=56: the maximum link rate in kbps
