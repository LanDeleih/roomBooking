package api

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/lanDeleih/roomBooking/internal/roomRepository"
	log "github.com/sirupsen/logrus"
)

//Api realised server methods
type Api interface {
	Start() error
	Shutdown(ctx context.Context) error
}

type Router struct {
	Server         *echo.Echo
	Logger         *log.Logger
	RoomRepository roomRepository.RoomRepository
}

func NewApi(logger *log.Logger) Api {
	e := echo.New()
	prom := prometheus.NewPrometheus("roomBooking", nil)
	prom.Use(e)
	rr := roomRepository.NewRoomRepository(logger)
	srv := &Router{
		Server:         e,
		Logger:         logger,
		RoomRepository: rr,
	}
	srv.configureRoutes()
	return srv
}
// Start will start http server
func (r *Router) Start() error {
	return r.Server.Start(":8080")
}
// Shutdown produce graceful shutdown for http server
func (r *Router) Shutdown(ctx context.Context) error {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return r.Server.Shutdown(ctx)
}
