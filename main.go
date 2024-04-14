package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/SnrMatt/bugdet.io/api"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	/*
		Database connection
			Note: This should be cleaned up later.
	*/

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found.")
	}

	uri := os.Getenv("DATABASE_URI")
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opt := options.Client().ApplyURI(uri).SetServerAPIOptions(serverApi)

	client, err := mongo.Connect(context.TODO(), opt)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	if err := client.Database("budget-io").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	/*
		Database connection
			End
	*/
	listenAddr := flag.String("listenAddr", ":3001", "Server address")
	flag.Parse()

	s := api.NewServer(*listenAddr)

	fmt.Println("Server starting on http://localhost:3001 ...")
	log.Fatal(s.Start())
}
