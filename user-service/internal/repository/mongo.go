package repository

import (
	"context"
	"time"

	"github.com/Cypher042/PaaS/user-service/internal/user"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepo struct {
	Collection *mongo.Collection
}

func (r *UserRepo) Create(u *user.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.Collection.InsertOne(ctx, u)
	return err
}

func (r *UserRepo) FindUserByID(id uuid.UUID) (*user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var u user.User

	err := r.Collection.FindOne(ctx, bson.M{"id": id}).Decode(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}


func (r *UserRepo) FindUserByUsername(username string) (*user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var u user.User

	err := r.Collection.FindOne(ctx, bson.M{"username": username}).Decode(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}