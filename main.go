package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/nats-io/go-nats"
)

// Application info
const (
	version = "0.0.1"
	appName = "chatstatz-webhooks"
)

// Default CLI options
var (
	host        = "0.0.0.0"
	port        = 5994
	mqHost      = "nats://0.0.0.0:4222"
	mqQueue     = "webhooks"
	verbose     = false
	showVersion = false
)

var usageHelp = fmt.Sprintf(`
Usage: %s [options]

Options:
  -h <host>          the host to expose this service on (default: %s)
  -p <port>          the port to run this server on (default: %d)
  -mqh <mq_host>     the message queue host server (default: %s)
  -mqq <mq_queue>    the name of the queue to publish messages to (default: %s)
  -v                 enable verbose mode
  --version          print the %s version
  --help             show help
`, appName, host, port, mqHost, mqQueue, appName)

func init() {
	flag.StringVar(&host, "h", host, "the host to expose this service on")
	flag.IntVar(&port, "p", port, "the port to run this server on")
	flag.StringVar(&mqHost, "mqh", mqHost, "the message queue host server.")
	flag.StringVar(&mqQueue, "mqq", mqQueue, "the name of the queue to publish messages to.")
	flag.BoolVar(&verbose, "v", verbose, "enable verbose mode")
	flag.BoolVar(&showVersion, "version", false, fmt.Sprintf("show %s version.", appName))

	flag.Usage = usage
	flag.Parse()
}

var service ServiceInterface

func main() {
	if showVersion {
		printVersion()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	server := newServer(host, port)
	producer, err := newProducer(mqHost, mqQueue, newNatsProducer)
	if err != nil {
		lfatalf("Failed to estasblish a connection for producer: %s", err.Error())
	}

	service = newService(&ServiceOptions{
		Server:   server,
		Producer: producer,
	})

	go func() {
		linfof("Starting %s server on %s", appName, server.Addr)
		lfatal(service.Start())
	}()

	<-quit

	linfo("SIGINT received... Stopping server.")

	// Cleanup before quitting
	service.Stop()
}

func newNatsProducer(host, queue string) (*Producer, error) {
	conn, err := nats.Connect(host, nats.Timeout(3*time.Second))
	if err != nil {
		return nil, err
	}

	return &Producer{
		conn:  conn,
		host:  host,
		queue: queue,
	}, nil
}

func usage() {
	fmt.Println(usageHelp)
	os.Exit(0)
}

func printVersion() {
	fmt.Printf("%s v%s\n", appName, version)
	os.Exit(0)
}
