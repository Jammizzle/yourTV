package remote

import (
	"net"
	"net/http"
	"time"
)

var client *http.Client

func init() {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	client = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
}
