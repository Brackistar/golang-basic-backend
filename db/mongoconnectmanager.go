package db

import (
	"context"
	"fmt"
	"log"

	"github.com/Brackistar/golang-basic-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnectManager struct {
	client *mongo.Client
	dbName string
}

func NewMongoConnectManager() *MongoConnectManager {
	return &MongoConnectManager{}
}

func (i *MongoConnectManager) GetDbName() string {
	return i.dbName
}

func (i *MongoConnectManager) Connect(ctx context.Context) error {
	user := ctx.Value(models.Key("user")).(string)
	pswrd := ctx.Value(models.Key("pswrd")).(string)
	host := ctx.Value(models.Key("host")).(string)
	connStr := fmt.Sprintf(mongoDbConnectionString, user, pswrd, host)

	var clientOptions = options.Client().ApplyURI(connStr)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Printf(errorMongoConnFailMsg, err.Error())
		return err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Printf(errorMongoConnFailMsg, err.Error())
		return err
	}

	log.Print("Connection stablished with MongoDB")

	i.client = client

	i.dbName = ctx.Value(models.Key("db")).(string)
	return nil
}

// Checks connection with the Database, returns true in case connection can be stablished
func (i *MongoConnectManager) IsConnected() bool {
	err := i.client.Ping(context.TODO(), nil)

	return err == nil
}
