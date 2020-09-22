package main

import (
	"context"
	"github.com/lanDeleih/roomBooking/internal/api"
	log "github.com/sirupsen/logrus"
)

func main() {
	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	s := api.NewApi(logger)
	ctx := context.Background()

	go func() {
		if err := s.Start(); err != nil {
			logger.Info("Shutting down the server: ", err)
		}
	}()
	logger.Info("prometheus metrics available on /metrics")
	if err := s.Shutdown(ctx); err != nil {
		logger.Fatal(err)
	}
}
