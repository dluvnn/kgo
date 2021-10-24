package kgo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ConnectDB connects to MongoDB server and returns the client object
func ConnectDB(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	clientOptions := options.Client().ApplyURI(uri)
	var err error
	c, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	return c, c.Ping(ctx, readpref.Primary())
}

// RunMongoSession calls function fx in a MongoDB session
func RunMongoSession(c *mongo.Client, ctx context.Context, fx func(mongo.SessionContext) (interface{}, error)) (interface{}, error) {
	session, err := c.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)
	return session.WithTransaction(ctx, fx)
}
