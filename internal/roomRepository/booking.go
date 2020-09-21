package roomRepository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

const (
	layoutTime      = "01/02/2006 3:04 PM"
	minimalDuration = 30 * time.Minute
)

type BookingRepository interface {
	CreateReservation(room Appointment) error
	GetAllReservation() (interface{}, error)
}

// CreateReservation will create new reservation if not found
func (m *MongoDB) CreateReservation(appointment Appointment) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	startDate, err := time.Parse(layoutTime, appointment.StartDate)
	if err != nil {
		m.Logger.Errorf("error parsing start date for reservation: %s", err)
	}
	if time.Now().Unix() > startDate.Unix() {
		return errors.New("date is not valid")
	}
	d, err := time.ParseDuration(appointment.Duration)
	if err != nil {
		m.Logger.Errorf("error in duration parse", err)
	}
	if minimalDuration.Minutes() > d.Minutes() {
		return errors.New("duration less than 30m")
	}
	endReservation := startDate.Add(d)
	roomExist := m.FindRoom(ctx, appointment.Room)
	if !roomExist {
		return errors.New("room doesn't exist")
	}
	find := m.FindReservation(ctx, appointment, endReservation, startDate)
	if find {
		return errors.New("appointment found")
	}
	document := bson.M{
		"customer":  appointment.Customer,
		"startDate": startDate,
		"endDate":   endReservation,
		"room":      appointment.Room,
	}
	_, err = m.DB.Collection("booking").InsertOne(ctx, document)
	if err != nil {
		return err
	}
	return nil
}

// GetAllReservation return all reservation from database
func (m *MongoDB) GetAllReservation() (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	var a MongoAppointment
	var result []MongoAppointment
	curr, err := m.DB.Collection("booking").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for curr.TryNext(ctx) {
		if err := curr.Decode(&a); err != nil {
			return nil, err
		}
		result = append(result, a)
	}
	return result, nil
}

// FindReservation return true if reservation was found
func (m *MongoDB) FindReservation(ctx context.Context, appointment Appointment, endReservation, startReservation time.Time) bool {
	filter := bson.M{
		"room":      appointment.Room,
		"startDate": bson.M{"$lte": startReservation},
		"endDate":   bson.M{"$lte": endReservation},
	}
	var a MongoAppointment
	curr, err := m.DB.Collection("booking").Find(ctx, filter)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false
		}
		m.Logger.Errorf("failed to decode result: %s", err)
		return true
	}
	for curr.TryNext(ctx) {
		if err := curr.Decode(&a); err != nil {
			m.Logger.Error(err)
		}
		if isTimeBetween(a.StartDate, a.EndDate, startReservation) {
			return true
		}
		continue

	}
	return false
}

func isTimeBetween(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}
