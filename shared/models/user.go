package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name,omitempty"`
	LastName  string             `bson:"lastname" json:"lastname,omitempty"`
	BirthDate time.Time          `bson:"birthDate" json:"birthDate,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Pass      string             `bson:"pass" json:"pass,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Bio       string             `bson:"bio" json:"bio,omitempty"`
	Local     string             `bson:"local" json:"local,omitempty"`
	URLs      []string           `bson:"urls" json:"urls,omitempty"`
}
