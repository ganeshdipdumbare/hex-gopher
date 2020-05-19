package mongodb

import (
	"context"
	"errors"
	"hex-gopher/app"
	"time"

	errs "github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	client           *mongo.Client
	gopherDb         string
	gopherCollection string
}

func getMongoClient(uri string) (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errs.Wrap(err, "failed to create client for mongo")
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, errs.Wrap(err, "failed to ping with client for mongo")
	}
	return client, nil
}

func NewMongoDB(uri string, db string, collection string) (*MongoDB, error) {
	cl, err := getMongoClient(uri)
	if err != nil {
		return nil, errs.Wrap(err, "failed to get client for redis")
	}

	return &MongoDB{
		client:           cl,
		gopherDb:         db,
		gopherCollection: collection,
	}, nil
}

func (r *MongoDB) SaveGopher(g *app.Gopher) (string, error) {
	if g == nil {
		return "", errs.Wrap(errors.New("nil gopher passed"), "nil gopher passed")
	}

	_, err := r.client.Database(r.gopherDb).Collection(r.gopherCollection).InsertOne(context.Background(), g)
	if err != nil {
		return "", errs.Wrap(err, "unable to save the gopher to mongo")
	}
	return g.Id, nil
}

func (r *MongoDB) GetGopher(id string) (*app.Gopher, error) {
	gopher := &app.Gopher{}
	err := r.client.Database(r.gopherDb).Collection(r.gopherCollection).FindOne(context.Background(), primitive.M{
		"id": id,
	}).Decode(&gopher)
	if err != nil {
		return nil, errs.Wrap(err, "unable to get the gopher from redis")
	}
	return gopher, nil
}
