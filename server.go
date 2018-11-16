package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/urfave/negroni"
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
		Handler:      middlewareWrapper(mux),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func middlewareWrapper(mux *http.ServeMux) *negroni.Negroni {
	n := negroni.New()
	recovery := negroni.NewRecovery()
	recovery.Logger = log.New(os.Stderr, "", log.LstdFlags)
	recovery.Formatter = &PanicFormatter{}
	n.Use(recovery)
	n.UseHandler(mux)
	return n
}

// PanicFormatter implements negroni.PanicFormatter.
type PanicFormatter struct{}

// FormatPanicError formats the response for a given panic.
func (pf *PanicFormatter) FormatPanicError(rw http.ResponseWriter, r *http.Request, infos *negroni.PanicInformation) {
	rw.Header().Set("Content-Type", "application/json") // See https://github.com/urfave/negroni/issues/241
	rw.Write([]byte(`{"success":"false","message":"Internal Server Error"}`))
}
