package main

import (
	"github.com/lanDeleih/roomBooking/internal/api"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	logger := log.New()
	s := api.NewApi(logger)
	if err := s.Start(); err != nil {
		logger.Fatalf("Failed to start server: %s", err)
	}
}

func newDB() (error) {

	return nil
}