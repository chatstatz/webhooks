package main

import (
	"flag"
	"fmt"
	"net"
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
Usage: chatstatz-webhooks [flags]

Options:
  --version     print the version
  --list-env    list expected environment variables and their default values
  --help        show help
`

func main() {
	logger := logger.New(EnvLogLevel, os.Stderr)
	logger.Debugf("Environment Variables: %s", strings.Join(os.Environ(), " "))

	var showVersion bool
	var listEnvVars bool

	flag.BoolVar(&showVersion, "version", false, "show webhooks version")
	flag.BoolVar(&listEnvVars, "list-env", false, "list expected environment variables and their default values")

	flag.Usage = usage
	flag.Parse()

	if showVersion {
		printVersion()
	}

	if listEnvVars {
		printEnvVars()
	}

	logger.Debugf("Starting webhooks server on %s:%s", EnvWebhooksHost, EnvWebhooksPort)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	conn, err := natsc.Connect(net.JoinHostPort(EnvNatsHost, EnvNatsPort), natsc.Timeout(3*time.Second))
	if err != nil {
		logger.Fatalf("Failed to estasblish a connection to NATS: %s", err.Error())
	}

	producer := nats.NewProducer(conn, EnvNatsQueue)
	ctx := context.NewContext(logger, producer)

	httpServer := http.NewServer(ctx, EnvWebhooksHost, EnvWebhooksPort)
	service := internal.NewWebhooksService(httpServer, producer)

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

func printVersion() {
	fmt.Printf("%s v%s\n", appName, version)
	os.Exit(0)
}

func printEnvVars() {
	fmt.Printf(`WEBHOOKS_HOST	The host that the  webhooks server should run on (default: 127.0.0.1)
WEBHOOKS_PORT	The port that the webhooks server should be served on (default: 8080)
NATS_HOST	The NATS host address for NATS clients to connect to (default: 0.0.0.0)
NATS_PORT	The NATS port for NATS clients to connect on (default: 4222)
NATS_QUEUE	The NATS queue for which to publish messages to (default: twitch_channels)
LOG_LEVEL	The log level to start logging from (default: info)
`)
	os.Exit(0)
}
