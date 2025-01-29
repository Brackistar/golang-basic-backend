package db

import (
	"context"
	"fmt"
	"log"

	"github.com/Brackistar/golang-basic-backend/shared/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDataOrigin struct {
	client *mongo.Client
	DbName string
}

func CreateMongoDataOrigin(c *mongo.Client) *MongoDataOrigin {
	return &MongoDataOrigin{
		client: c,
	}
}

func (i *MongoDataOrigin) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return i.client.Ping(ctx, rp)
}

func (i *MongoDataOrigin) CreateRecord(source any, args ...any) (any, bool, error) {
	log.Printf("Creating record on source: %s, for values %s", source, args)

	colName, ok := source.(string)

	if !ok {
		return nil, false, fmt.Errorf(invColNameMsg, source)
	}

	collection := i.client.Database(i.DbName).Collection(colName)

	result, err := collection.InsertOne(context.TODO(), args[0])

	if err != nil {
		return nil, false, err
	}

	id, _ := result.InsertedID.(primitive.ObjectID)

	return id.String(), true, nil
}

// Search for a single record using a single
func (i *MongoDataOrigin) GetRecord(source any, args ...any) (any, error) {

	log.Printf("Retrieving record on source: %s, for values %v", source, args)

	colName, ok := source.(string)

	if !ok {
		return nil, fmt.Errorf(invColNameMsg, source)
	}

	log.Printf("Retrieving collection by name: %s", colName)
	collection := i.client.Database(i.DbName).Collection(colName)

	log.Print("Collection found")

	condition := bson.M{
		args[0].(string): args[1].(string),
	}

	var result models.User

	err := collection.FindOne(context.TODO(), condition).Decode(&result)

	if err != nil {
		log.Print(err.Error())

		return nil, err
	}

	log.Printf("Record found: %v", result)

	return result, nil
}

func (i *MongoDataOrigin) UpdateRecord(val any, args ...any) error {
	return nil
}
func (i *MongoDataOrigin) DeleteRecord(id any, args ...any) error {
	return nil
}
