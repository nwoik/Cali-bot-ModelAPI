package clan_test

import (
	"context"
	"log"
	"testing"

	c "github.com/nwoik/calibotapi/model/clan"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://mongo:prcifwVyTGyjhszMBDAaDntdsxSEJJLi@viaduct.proxy.rlwy.net:58839/?tlsCertificateKeyFilePassword=prcifwVyTGyjhszMBDAaDntdsxSEJJLi"))

	if err != nil {
		log.Fatal("error connecting to db", err)
	}

	log.Println("successfully connected")

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed")
	}

	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := NewMongoClient()

	defer mongoTestClient.Disconnect(context.Background())

	collection := mongoTestClient.Database("calibot").Collection("clan_test")

	clanRepo := c.NewClanRepo(collection)

	t.Run("Insert first clan", func(t *testing.T) {
		clan := c.CreateClan("my clan", "124312424", "124324253456")
		result, err := clanRepo.InsertClan(clan)

		if err != nil {
			t.Fatal("Failed insertion", err)
		}

		t.Log("Insert successful", result)
	})

	t.Run("Insert second clan", func(t *testing.T) {
		clan := c.CreateClan("my clan2", "124344344", "234323463456")
		result, err := clanRepo.InsertClan(clan)

		if err != nil {
			t.Fatal("Failed insertion", err)
		}

		t.Log("Insert successful", result)
	})

	t.Run("Get 2nd clan", func(t *testing.T) {
		result, err := clanRepo.GetClanByID("124344344")

		if err != nil {
			t.Fatal("Failed to get clan", err)
		}

		t.Log("clan:", result.Name)
	})

	t.Run("Get all clans", func(t *testing.T) {
		results, err := clanRepo.GetAllClans()

		if err != nil {
			t.Fatal("Failed to get clans", err)
		}

		t.Log("clan:", results)
	})

	t.Run("Update 2nd clan", func(t *testing.T) {
		clan, err := clanRepo.GetClanByID("124344344")

		if err != nil {
			t.Fatal("Failed to get clan", err)
		}

		clan.LeaderID = "deeznuts"

		i, err := clanRepo.UpdateClan(clan)

		if err != nil {
			log.Fatal("update failed", err)
		}

		t.Log("number of clans updated:", i)
	})

	t.Run("Update 1st clan", func(t *testing.T) {
		clan, err := clanRepo.GetClanByID("124324253456")

		if err != nil {
			t.Fatal("Failed to get clan", err)
		}

		clan.LeaderID = "nuts"

		i, err := clanRepo.UpdateClan(clan)

		if err != nil {
			log.Fatal("update failed", err)
		}

		t.Log("number of clans updated:", i)
	})
}
