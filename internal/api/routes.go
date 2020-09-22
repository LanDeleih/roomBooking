package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lanDeleih/roomBooking/internal/roomRepository"
	"strconv"
)

func (r *Router) configureRoutes() {
	r.Server.Use(middleware.Recover())
	r.Server.Use(middleware.BodyLimit("2M"))
	v1 := r.Server.Group("/api/v1")
	room := v1.Group("/room")
	{
		room.POST("/create", func(c echo.Context) error {
			var cr roomRepository.Room
			if err := c.Bind(&cr); err != nil {
				return c.JSON(400, err)
			}
			if cr.Number == 0 {
				return c.JSON(400, map[string]string{"error": "room is not specified"})
			}
			if err := r.RoomRepository.CreateRoom(cr.Number); err != nil {
				return c.JSON(500, map[string]string{"error": err.Error()})
			}
			return c.JSON(200, map[string]string{"message": "room created"})
		})
		room.POST("/delete/:id", func(c echo.Context) error {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				return c.JSON(500, map[string]string{"error": err.Error()})
			}
			if err := r.RoomRepository.DeleteRoom(int32(id)); err != nil {
				return c.JSON(500, map[string]string{"error": err.Error()})
			}
			return c.JSON(200, map[string]string{"message": "room deleted"})
		})
		room.GET("/", func(c echo.Context) error {
			rooms, err := r.RoomRepository.GetRooms()
			if err != nil {
				return c.JSON(500, map[string]string{"error": err.Error()})
			}
			return c.JSON(200, rooms)
		})
	}
	booking := v1.Group("/booking")
	{
		booking.POST("/create", func(c echo.Context) error {
			var ra roomRepository.Appointment
			if err := c.Bind(&ra); err != nil {
				return c.JSON(400, err)
			}
			if ra.Room == 0 {
				return c.JSON(400, map[string]string{"error": "room is not specified"})
			}
			if err := r.RoomRepository.CreateReservation(ra); err != nil {
				return c.JSON(500, map[string]string{"error": err.Error()})
			}
			return c.JSON(200, map[string]string{"message": "reservation created"})
		})
		booking.GET("/", func(c echo.Context) error {
			ar, err := r.RoomRepository.GetAllReservation()
			if err != nil {
				return c.JSON(500, map[string]string{"error": err.Error()})
			}
			return c.JSON(200, ar)
		})
		booking.GET("/:id", func(c echo.Context) error {
			return nil
		})
	}

}

