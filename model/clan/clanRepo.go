package clan

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClanRepo struct {
	MongoCollection *mongo.Collection
}

func NewClanRepo(collection *mongo.Collection) *ClanRepo {
	return &ClanRepo{MongoCollection: collection}
}

func (clanRepo *ClanRepo) Delete(id string) (int64, error) {
	result, err := clanRepo.MongoCollection.DeleteOne(context.Background(),
		bson.D{{Key: "clanid", Value: id}})

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func (clanRepo *ClanRepo) DeleteAll() (int64, error) {
	result, err := clanRepo.MongoCollection.DeleteMany(context.Background(),
		bson.D{})

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func (clanRepo *ClanRepo) Get(id string) (*Clan, error) {
	var clan Clan

	err := clanRepo.MongoCollection.FindOne(context.Background(), bson.D{{Key: "clanid", Value: id}}).Decode(&clan)
	if err != nil {
		err = clanRepo.MongoCollection.FindOne(context.Background(), bson.D{{Key: "guildid", Value: id}}).Decode(&clan)
		if err != nil {
			return nil, err
		}
	}

	return &clan, nil
}

func (clanRepo *ClanRepo) GetAll() ([]Clan, error) {
	results, err := clanRepo.MongoCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var clans []Clan
	err = results.All(context.Background(), &clans)
	if err != nil {
		return nil, fmt.Errorf("results decode error %s", err.Error())
	}

	return clans, err
}

func (clanRepo *ClanRepo) Insert(clan *Clan) (interface{}, error) {
	result, err := clanRepo.MongoCollection.InsertOne(context.Background(), clan)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (clanRepo *ClanRepo) Update(clan *Clan) (int64, error) {
	result, err := clanRepo.MongoCollection.UpdateOne(context.Background(),
		bson.D{{Key: "clanid", Value: clan.ClanID}},
		bson.D{{Key: "$set", Value: clan}})

	if err != nil {
		return 0, err
	}

	return result.UpsertedCount, nil
}
