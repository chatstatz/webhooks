package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/chatstatz/logger"
	"github.com/chatstatz/webhooks/internal"
	"github.com/chatstatz/webhooks/internal/context"
	"github.com/chatstatz/webhooks/internal/http"
	"github.com/chatstatz/webhooks/internal/nats"
	"github.com/chatstatz/webhooks/tools"

	natsc "github.com/nats-io/nats.go"
)

// Application info
const (
	version = "0.1.0"
	appName = "chatstatz-webhooks"
)

// Environment variable defaults
var (
	WebhooksHostDefault = "127.0.0.0"
	WebhooksPortDefault = "8080"
	NatsHostDefault     = "0.0.0.0"
	NatsPortDefault     = "4222"
	NatsQueueDefault    = "twitch_channels"
	LogLevelDefault     = "info"
)

// Environment variables
var (
	EnvWebhooksHost = tools.GetEnv("WEBHOOKS_HOST", WebhooksHostDefault)
	EnvWebhooksPort = tools.GetEnv("WEBHOOKS_PORT", WebhooksPortDefault)
	EnvNatsHost     = tools.GetEnv("NATS_HOST", NatsHostDefault)
	EnvNatsPort     = tools.GetEnv("NATS_PORT", NatsPortDefault)
	EnvNatsQueue    = tools.GetEnv("NATS_QUEUE", NatsQueueDefault)
	EnvLogLevel     = tools.GetEnv("LOG_LEVEL", LogLevelDefault)
)

var usageHelp = `
Usage: %s [options]

Options:
  -v                    print the %s version
  --help                show help
`

func main() {
	logger := logger.New(EnvLogLevel, os.Stderr)
	logger.Debugf("Environment Variables: %s", strings.Join(os.Environ(), " "))

	var printVersion bool
	flag.BoolVar(&printVersion, "v", false, "show webhooks version.")
	flag.Usage = usage
	flag.Parse()

	if printVersion {
		showVersion()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	conn, err := natsc.Connect(NatsHostDefault, natsc.Timeout(3*time.Second))
	if err != nil {
		logger.Fatalf("Failed to estasblish a connection to NATS: %s", err.Error())
	}

	producer := nats.NewProducer(conn, NatsPortDefault)
	ctx := context.NewContext(logger, producer)

	httpServer := http.NewServer(ctx, EnvWebhooksHost, EnvWebhooksPort)
	service := internal.NewService(httpServer, producer)

	go func() {
		logger.Infof("Starting %s server on %s", appName, httpServer.Addr)
		logger.Fatal(service.Start())
	}()

	<-quit

	logger.Infof("SIGINT received... Stopping server.")

	// Cleanup before quitting
	service.Stop()
}

func usage() {
	fmt.Println(usageHelp)
	os.Exit(0)
}

func showVersion() {
	fmt.Printf("%s v%s\n", appName, version)
	os.Exit(0)
}
