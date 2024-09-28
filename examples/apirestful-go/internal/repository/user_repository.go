package repository

import (
	"apirestful-go/internal/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(id primitive.ObjectID) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(id primitive.ObjectID, user models.User) (models.User, error)
	Delete(id primitive.ObjectID) error
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}

func (r *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetByID(id primitive.ObjectID) (models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	return user, err
}

func (r *userRepository) Create(user models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, err := r.collection.InsertOne(ctx, user)

	if err != nil {
		return models.User{}, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (r *userRepository) Update(id primitive.ObjectID, user models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": user})
	if err != nil {
		return models.User{}, err
	}
	user.ID = id
	return user, nil
}

func (r *userRepository) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
