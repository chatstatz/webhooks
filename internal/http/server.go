package http

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/chatstatz/webhooks/internal/context"
	"github.com/urfave/negroni"
)

// Server ...
type Server struct {
	ctx *context.Context
}

// NewServer creates a new HTTP server.
func NewServer(ctx *context.Context, host, port string) *http.Server {
	server := &Server{ctx}
	mux := http.NewServeMux()

	mux.HandleFunc("/healhtz", server.healthCheckHandler)
	mux.HandleFunc("/twitch/webhooks", server.handleWebhookPostRequest)

	return &http.Server{
		Addr:         net.JoinHostPort(host, port),
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

	// TODO: log panic error and stack

	rw.Header().Set("Content-Type", "application/json") // See https://github.com/urfave/negroni/issues/241
	rw.Write([]byte(`{"success":"false","message":"Internal Server Error"}`))
}
