package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Username   string             `bson:"username"`
	Name       string             `bson:"name"`
	Email      string             `bson:"email"`
	Age        int                `bson:"age"`
	DNI        string             `bson:"dni"`
	Phone      string             `bson:"phone"`
	Country    string             `bson:"country"`
	State      string             `bson:"state"`
	City       string             `bson:"city"`
	Address    string             `bson:"address"`
	PostalCode string             `bson:"postal_code"`
	Password   string             `bson:"password"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
	DeletedAt  time.Time          `bson:"deleted_at"`
}
