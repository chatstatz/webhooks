package main

import (
	"net"
	"net/http"
	"strconv"
	"time"
)

// ServerInterface ...
type ServerInterface interface {
	ListenAndServe() error
}

type route struct {
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
}

var routes = []route{
	route{"/health-check", healthCheckHandler},
	route{"/twitch/webhooks", twitchWebhookHandler},
}

func newServer(host string, port int) *http.Server {
	mux := http.NewServeMux()

	for _, route := range routes {
		mux.HandleFunc(route.Path, route.Handler)
	}

	return &http.Server{
		Addr:         net.JoinHostPort(host, strconv.Itoa(port)),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
