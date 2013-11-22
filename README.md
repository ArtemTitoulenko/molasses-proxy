# Molasses Proxy

Molasses Proxy waits longer and longer to return a response for a list of blocked
hosts. It's useful for slow-banning computers on a network and discouraging the
use of certain websites.

## Installation
First, install the binary.

    $ go get github.com/ArtemTitoulenko/molasses-proxy

Then in some directory where you plan on running the proxy, create a
`blocked_hosts` file and list websites you would want to block on separate
lines. An example config is supplied.

## Running

    $ molasses-proxy

## Getting Help

    $ molasses-proxy --help
      usage:

          molasses-proxy [--port=8080] [--delayms=500]

      -delayms=500: increase the delay by this many milliseconds per request
      -help=false: print this help message
      -port=8080: the port to listen for requests on
