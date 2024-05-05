package clan_test

import (
	"context"
	"log"
	"os"
	"testing"

	c "github.com/nwoik/calibotapi/model/clan"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoClient() *mongo.Client {
	pswd := os.Getenv("MONGO_PASS")
	mongoClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://mongo:"+pswd+"@viaduct.proxy.rlwy.net:58839/?tlsCertificateKeyFilePassword="+pswd))

	if err != nil {
		log.Fatal("error connecting to db", err)
	}

	log.Println("successfully connected")

	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed")
	}

	return mongoClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := NewMongoClient()

	defer mongoTestClient.Disconnect(context.Background())

	collection := mongoTestClient.Database("calibot").Collection("clan_test")

	clanRepo := c.NewClanRepo(collection)

	t.Run("Insert first clan", func(t *testing.T) {
		clan := c.CreateClan("my clan", "124312424", "124324253456")
		result, err := clanRepo.Insert(clan)

		if err != nil {
			t.Fatal("Failed insertion", err)
		}

		t.Log("Insert successful", result)
	})

	t.Run("Insert second clan", func(t *testing.T) {
		clan := c.CreateClan("my clan2", "124344344", "234323463456")
		result, err := clanRepo.Insert(clan)

		if err != nil {
			t.Fatal("Failed insertion", err)
		}

		t.Log("Insert successful", result)
	})

	t.Run("Get 2nd clan", func(t *testing.T) {
		result, err := clanRepo.Get("124344344")

		if err != nil {
			t.Fatal("Failed to get clan", err)
		}

		t.Log("clan:", result.Name)
	})

	t.Run("Get all clans", func(t *testing.T) {
		results, err := clanRepo.GetAll(bson.E{Key: "clanid", Value: "124312424"})

		if err != nil {
			t.Fatal("Failed to get clans", err)
		}

		t.Log("clan:", results)
	})

	t.Run("Update 2nd clan", func(t *testing.T) {
		clan, err := clanRepo.Get("124344344")

		if err != nil {
			t.Fatal("Failed to get clan", err)
		}

		clan.LeaderID = "deeznuts"

		i, err := clanRepo.Update(clan)

		if err != nil {
			log.Fatal("update failed", err)
		}

		t.Log("number of clans updated:", i)
	})

	t.Run("Update 1st clan", func(t *testing.T) {
		clan, err := clanRepo.Get("124324253456")

		if err != nil {
			t.Fatal("Failed to get clan", err)
		}

		clan.LeaderID = "nuts"

		i, err := clanRepo.Update(clan)

		if err != nil {
			log.Fatal("update failed", err)
		}

		t.Log("number of clans updated:", i)
	})

	t.Run("Delete all", func(t *testing.T) {
		i, err := clanRepo.DeleteAll()

		if err != nil {
			log.Fatal("deletions failed", err)
		}

		t.Log("number of clans deleted:", i)
	})
}
