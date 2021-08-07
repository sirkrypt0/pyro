package main

import (
	"errors"
	"flag"
	"github.com/sirkrypt0/pyro/internal/agent"
	"github.com/sirkrypt0/pyro/pkg/logging"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := logging.GetLogger("pyro-agent")

	bindAddress := flag.String("bind", "127.0.0.1:3000", "Address to bind to")
	flag.Parse()

	srv, err := agent.NewGRPCServer()
	if err != nil {
		log.WithError(err).Fatal("Error creating new GRPC server!")
	}

	go func() {
		l, err := net.Listen("tcp", *bindAddress)
		if err != nil {
			log.WithError(err).Fatal("Error during listening")
		}
		log.WithField("address", *bindAddress).Info("Started listening ...")
		if err = srv.Serve(l); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.WithError(err).Info("Server closed")
			} else {
				log.WithError(err).Fatal("Error during serving")
			}
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Info("Received SIGINT, shutting down ...")
	srv.GracefulStop()
}
