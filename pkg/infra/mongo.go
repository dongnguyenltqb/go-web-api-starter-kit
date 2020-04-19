package infra

import (
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func GetDB() *mongo.Database {
	if db == nil {
		panic(errors.New("setup mongodb first..."))
	}
	return db
}

func setupMongo() {
	conf := viper.GetViper()
	mongoURI := conf.GetString(`mongodb.url`)
	dbName := conf.GetString(`mongodb.db`)
	if mongoURI == "" {
		panic(errors.New("mongo uri is wrong, take a look"))
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}
	if err = client.Connect(context.Background()); err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err = client.Ping(ctx, nil); err != nil {
		panic(err)
	}
	logrus.Info("Setup MongoDB sucessfully...")
	db = client.Database(dbName)
}

func testWrite() {
	_, err := GetDB().Collection("test_write").InsertOne(context.Background(), bson.M{
		"time": time.Now().String(),
	})
	if err != nil {
		panic(err)
	}
}
