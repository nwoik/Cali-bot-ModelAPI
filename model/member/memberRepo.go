package member

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MemberRepo struct {
	MongoCollection *mongo.Collection
}

func NewMemberRepo(collection *mongo.Collection) *MemberRepo {
	return &MemberRepo{MongoCollection: collection}
}

func (memberRepo *MemberRepo) Delete(id string) (int64, error) {
	result, err := memberRepo.MongoCollection.DeleteOne(context.Background(),
		bson.D{{Key: "userid", Value: id}})

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func (memberRepo *MemberRepo) DeleteAll() (int64, error) {
	result, err := memberRepo.MongoCollection.DeleteMany(context.Background(),
		bson.E{})

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func (memberRepo *MemberRepo) Get(id string) (*Member, error) {
	var member Member

	err := memberRepo.MongoCollection.FindOne(context.Background(), bson.D{{Key: "userid", Value: id}}).Decode(&member)
	if err != nil {
		err = memberRepo.MongoCollection.FindOne(context.Background(), bson.D{{Key: "igid", Value: id}}).Decode(&member)
		if err != nil {
			return nil, err
		}
	}

	return &member, nil
}

func (memberRepo *MemberRepo) GetAll() ([]Member, error) {
	results, err := memberRepo.MongoCollection.Find(context.Background(), bson.E{})

	if err != nil {
		return nil, err
	}

	var members []Member
	err = results.All(context.Background(), &members)
	if err != nil {
		return nil, fmt.Errorf("results decode error %s", err.Error())
	}

	return members, err
}

func (memberRepo *MemberRepo) Filter(predicates ...bson.E) ([]Member, error) {
	filter := bson.D{}

	for _, predicate := range predicates {
		filter = append(filter, predicate)
	}
	results, err := memberRepo.MongoCollection.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	var members []Member
	err = results.All(context.Background(), &members)
	if err != nil {
		return nil, fmt.Errorf("results decode error %s", err.Error())
	}

	return members, err
}

func (memberRepo *MemberRepo) Insert(member *Member) (interface{}, error) {
	result, err := memberRepo.MongoCollection.InsertOne(context.Background(), member)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (memberRepo *MemberRepo) Update(member *Member) (int64, error) {
	result, err := memberRepo.MongoCollection.UpdateOne(context.Background(),
		bson.D{{Key: "userid", Value: member.ClanID}},
		bson.D{{Key: "$set", Value: member}})

	if err != nil {
		return 0, err
	}

	return result.UpsertedCount, nil
}
