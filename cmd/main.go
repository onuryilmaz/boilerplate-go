package main

import (
	"flag"
	"os"

	"os/signal"
	"sync"
	"syscall"

	"github.com/Sirupsen/logrus"
	"github.com/onuryilmaz/boilerplate-go/pkg/commons"
	"github.com/onuryilmaz/boilerplate-go/pkg/server"
	"github.com/spf13/pflag"
)

var options commons.Options

func init() {
	pflag.StringVar(&options.ServerPort, "port", "8080", "Server port for listening REST calls")
	pflag.StringVar(&options.LogLevel, "log-level", "info", "Log level, options are panic, fatal, error, warning, info and debug")
}

func main() {

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	level, err := logrus.ParseLevel(options.LogLevel)
	if err == nil {
		logrus.SetLevel(level)
	} else {
		logrus.Fatal("Error during log level parse:", err)
	}

	sigs := make(chan os.Signal, 1)
	stop := make(chan struct{})
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	wg := &sync.WaitGroup{}

	webserver := server.NewREST(options)
	webserver.Start()

	<-sigs
	logrus.Warn("Shutting down...")
	webserver.Stop()

	close(stop)
	wg.Wait()
}
