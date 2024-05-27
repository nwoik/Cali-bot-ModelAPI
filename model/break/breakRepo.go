package breaks

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BreakRepo struct {
	MongoCollection *mongo.Collection
}

func NewBreakRepo(collection *mongo.Collection) *BreakRepo {
	return &BreakRepo{MongoCollection: collection}
}

func (breakRepo *BreakRepo) Delete(id string) (int64, error) {
	result, err := breakRepo.MongoCollection.DeleteOne(context.Background(),
		bson.D{{Key: "userid", Value: id}})

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func (breakRepo *BreakRepo) DeleteAll() (int64, error) {
	result, err := breakRepo.MongoCollection.DeleteMany(context.Background(),
		bson.D{})

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func (breakRepo *BreakRepo) Get(id string) (*Break, error) {
	var brk Break

	err := breakRepo.MongoCollection.FindOne(context.Background(), bson.D{{Key: "userid", Value: id}}).Decode(&brk)
	if err != nil {
		return nil, err
	}

	return &brk, nil
}

func (breakRepo *BreakRepo) GetAll() ([]*Break, error) {
	results, err := breakRepo.MongoCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var breaks []*Break
	err = results.All(context.Background(), &breaks)
	if err != nil {
		return nil, fmt.Errorf("results decode error %s", err.Error())
	}

	return breaks, err
}

func (breakRepo *BreakRepo) Filter(predicates ...bson.E) ([]*Break, error) {
	filter := bson.D{}

	for _, predicate := range predicates {
		filter = append(filter, predicate)
	}
	results, err := breakRepo.MongoCollection.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	var breaks []*Break
	err = results.All(context.Background(), &breaks)
	if err != nil {
		return nil, fmt.Errorf("results decode error %s", err.Error())
	}

	return breaks, err
}

func (breakRepo *BreakRepo) Insert(brk *Break) (interface{}, error) {
	result, err := breakRepo.MongoCollection.InsertOne(context.Background(), brk)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (breakRepo *BreakRepo) Update(brk *Break) (int64, error) {
	result, err := breakRepo.MongoCollection.UpdateOne(context.Background(),
		bson.D{{Key: "userid", Value: brk.UserID}},
		bson.D{{Key: "$set", Value: brk}})

	if err != nil {
		return 0, err
	}

	return result.UpsertedCount, nil
}
