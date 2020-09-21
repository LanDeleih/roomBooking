package roomRepository

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

type RoomRepository interface {
	Roomer
	BookingRepository
}

func NewRoomRepository(logger *log.Logger) RoomRepository {
	db, err := newDB()
	if err != nil {
		logger.Fatalf("failed to create database connection: %s", err)
	}
	if err := createCollection(db); err != nil {
		logger.Error(err)
	}
	return &MongoDB{
		DB:     db,
		Logger: logger,
	}
}

func newDB() (*mongo.Database, error) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return nil, errors.New("mongodb uri is not specified")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return client.Database("roomBooking"), nil
}

func createCollection(db *mongo.Database) error {
	if err := db.CreateCollection(context.TODO(), "rooms"); err != nil {
		log.Error(err)
	}
	if err := db.CreateCollection(context.TODO(), "booking"); err != nil {
		log.Error(err)
	}
	return nil
}
