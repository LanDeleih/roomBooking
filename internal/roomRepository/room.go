package roomRepository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Roomer interface {
	CreateRoom(room int32) error
	DeleteRoom(id int32) error
	GetRooms() (interface{}, error)
	FindRoom(ctx context.Context, room int32) bool
}

func (m *MongoDB) CreateRoom(room int32) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	document := m.FindRoom(ctx, room)
	if document {
		return errors.New("room already exist")
	}
	_, err := m.DB.Collection("rooms").InsertOne(ctx, bson.M{"number": room})
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) DeleteRoom(id int32) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	document := bson.M{"number": id}
	var r bson.M
	err := m.DB.Collection("rooms").FindOneAndDelete(ctx, document).Decode(&r)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return errors.New("room not found")
		}
	}
	return nil
}

func (m *MongoDB) GetRooms() (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	curr, err := m.DB.Collection("rooms").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var r Room
	var result []Room
	for curr.TryNext(ctx) {
		if err := curr.Decode(&r); err != nil {
			m.Logger.Infof("failed to find room: %s", err)
		}
		result = append(result, r)
	}
	return result, nil
}

func (m *MongoDB) FindRoom(ctx context.Context, room int32) bool {
	document := bson.M{"number": room}
	var r Room
	if err := m.DB.Collection("rooms").FindOne(ctx, document).Decode(&r); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false
		}
		m.Logger.Error(err)
	}
	if r.Number == room {
		return true
	}
	return false
}
