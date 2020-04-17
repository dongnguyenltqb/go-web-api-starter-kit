package infra

import (
	"context"
	"errors"
	"fmt"
	"time"

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
	fmt.Println("Setup MongoDB sucessfully...")
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
