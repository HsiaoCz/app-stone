package storage

import (
	"context"
	"errors"

	"github.com/HsiaoCz/app-stone/session/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStoreInter interface {
	CreateUser(context.Context, *types.Users) (*types.Users, error)
	GetUserByID(context.Context, primitive.ObjectID) (*types.Users, error)
	GetUserByEmail(context.Context, string) (*types.Users, error)
}

type UserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func UserStoreInit(client *mongo.Client, coll *mongo.Collection) *UserStore {
	return &UserStore{
		client: client,
		coll:   coll,
	}
}

func (u *UserStore) CreateUser(ctx context.Context, user *types.Users) (*types.Users, error) {
	filter := bson.M{
		"email":    user.Email,
		"username": user.Username,
	}
	cursor := u.coll.FindOne(ctx, filter)
	if cursor.Err() != mongo.ErrNoDocuments {
		return nil, errors.New("the record exists")
	}
	result, err := u.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, errors.New("db insert record error")
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (u *UserStore) GetUserByID(ctx context.Context, user_id primitive.ObjectID) (*types.Users, error) {
	var user types.Users
	filter := bson.D{
		{Key: "_id", Value: user_id},
	}
	if err := u.coll.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserStore) GetUserByEmail(ctx context.Context, email string) (*types.Users, error) {
	var user types.Users
	filter := bson.D{
		{Key: "email", Value: email},
	}
	if err := u.coll.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
