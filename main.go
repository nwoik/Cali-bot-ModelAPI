package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	m "github.com/nwoik/calibotapi/model/member"
)

// mongo "go.mongodb.org/mongo-driver/mongo"

func main() {
	url := os.Getenv("MONGO_URL")
	mongoClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal("error connecting to db", err)
	}

	defer mongoClient.Disconnect(context.Background())

	collection := mongoClient.Database("calibot").Collection("member")

	memberRepo := m.NewMemberRepo(collection)

	members, err := memberRepo.GetAll()

	if err != nil {
		log.Fatal("Failed to get members", err)
	}

	for _, member := range members {
		memberRepo.Update(member)
	}
}
