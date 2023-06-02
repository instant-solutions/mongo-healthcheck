package main

import (
	"context"
	"flag"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func main() {
	uri := flag.String("uri", "mongodb://localhost:27017", "MongoDB connection string")

	flag.Parse()

	if *uri == "" {
		fmt.Println("URI parameter missing")
		flag.Usage()
		os.Exit(1)
	}

	err := check(*uri)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully pinged MongoDB.")
}

func check(uri string) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)

	if err != nil {
		return err
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	var result bson.M

	if err := client.Database("admin").RunCommand(context.Background(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		return err
	}

	return nil
}
