package db

import (
	"context"
	"fmt"
	"log"
	"publish_it_everywhere/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db     *mongo.Database
	client *mongo.Client
	bg     = context.Background
)

// Initialize is used to initialize the db subsystem
func Initialize() {
	fmt.Println("inited DB")
	var err error
	client, err = mongo.Connect(bg(), options.Client().ApplyURI(config.DatabaseURL))
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	db = client.Database(config.DatabaseName)
}
