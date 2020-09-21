package roomRepository

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Room struct {
	Number int32 `json:"number"`
}

type Appointment struct {
	Customer  string    `json:"customer"`
	StartDate string    `json:"startDate"`
	Duration  string    `json:"duration,omitempty"`
	Room      int32     `json:"room"`
}

type MongoAppointment struct {
	Customer  string    `json:"customer"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Room      int32     `json:"room"`
}
type MongoDB struct {
	DB     *mongo.Database
	Logger *log.Logger
}
