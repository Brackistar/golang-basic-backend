package db

import (
	"context"
	"fmt"
	"log"

	"github.com/Brackistar/golang-basic-backend/interfaces"
	"github.com/Brackistar/golang-basic-backend/shared/constants"
	"github.com/Brackistar/golang-basic-backend/shared/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnectManager struct {
	dataOrigin MongoDataOrigin
}

func NewMongoConnectManager() *MongoConnectManager {
	return &MongoConnectManager{}
}

func (i *MongoConnectManager) GetDbName() string {
	return i.dataOrigin.DbName
}

func (i *MongoConnectManager) Connect(ctx context.Context) error {

	log.Println("Connecting with Mongo database")

	user := ctx.Value(constants.CtxKeyUser).(string)
	pswrd := ctx.Value(constants.CtxKeyPswd).(string)
	host := ctx.Value(constants.CtxKeyHost).(string)
	//database := ctx.Value(constants.CtxKeyDb).(string)
	connStr := fmt.Sprintf(mongoDbConnectionString, user, pswrd, host)

	log.Printf("Connection string: \"%s\"", connStr)

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

	i.dataOrigin = *CreateMongoDataOrigin(client)

	i.dataOrigin.DbName = utils.GetContextValue[string](&ctx, constants.CtxKeyDb)
	return nil
}

func (i *MongoConnectManager) GetDataOrigin() interfaces.DataOrigin {
	return &i.dataOrigin
}

// Checks connection with the Database, returns true in case connection can be stablished
func (i *MongoConnectManager) IsConnected() bool {
	err := i.dataOrigin.Ping(context.TODO(), nil)

	return err == nil
}
