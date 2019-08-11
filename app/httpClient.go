package app

import (
	"net"
	"net/http"
	"time"
)

const dialTimeout = 3 * time.Second
const tlsHandshakeTimeout = 3 * time.Second
const requestTimeout = time.Second * 10

// New - returns the singleton http app used for stapi requests
func createHttpClient() *http.Client {
	var netTransport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: dialTimeout,
		}).DialContext,
		TLSHandshakeTimeout: tlsHandshakeTimeout,
	}

	return &http.Client{
		Timeout:   requestTimeout,
		Transport: netTransport,
	}
}
