package member_test

import (
	"context"
	"log"
	"os"
	"testing"

	m "github.com/nwoik/calibotapi/model/member"

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

	collection := mongoTestClient.Database("calibot").Collection("member_test")

	memberRepo := m.NewMemberRepo(collection)

	t.Run("Insert first member", func(t *testing.T) {
		member := m.CreateMember("member 1", "m1", "124324256", "2948245235536345")
		result, err := memberRepo.Insert(member)

		if err != nil {
			t.Fatal("Failed insertion", err)
		}

		t.Log("Insert successful", result)
	})

	t.Run("Insert second member", func(t *testing.T) {
		member := m.CreateMember("member 2", "m2", "639834532", "3535636632353645")
		result, err := memberRepo.Insert(member)

		if err != nil {
			t.Fatal("Failed insertion", err)
		}

		t.Log("Insert successful", result)
	})

	t.Run("Get 2nd member", func(t *testing.T) {
		member, err := memberRepo.Get("3535636632353645")

		if err != nil {
			t.Fatal("Failed to get member", err)
		}

		t.Log("member:", member.Nick)
	})

	t.Run("Get all members", func(t *testing.T) {
		results, err := memberRepo.GetAll()

		if err != nil {
			t.Fatal("Failed to get members", err)
		}

		t.Log("member:", results)
	})

	t.Run("Update 2nd member", func(t *testing.T) {
		member, err := memberRepo.Get("3535636632353645")

		if err != nil {
			t.Fatal("Failed to get member", err)
		}

		member.Nick = "deeznuts"

		i, err := memberRepo.Update(member)

		if err != nil {
			log.Fatal("update failed", err)
		}

		t.Log("number of members updated:", i)
	})

	t.Run("Update 1st member", func(t *testing.T) {
		member, err := memberRepo.Get("124324256")

		if err != nil {
			t.Fatal("Failed to get member", err)
		}

		member.Nick = "nuts"

		i, err := memberRepo.Update(member)

		if err != nil {
			log.Fatal("update failed", err)
		}

		t.Log("number of members updated:", i)
	})

	t.Run("Delete all", func(t *testing.T) {
		i, err := memberRepo.DeleteAll()

		if err != nil {
			log.Fatal("deletions failed", err)
		}

		t.Log("number of members deleted:", i)
	})
}
