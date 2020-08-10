package data

import (
	"context"
	"fmt"

	"github.com/MurilloVaz/oratio/extensions/configuration"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataContext struct {
	*options.ClientOptions
	*mongo.Client
	connected bool
}

func NewDataContext() DataContext {
	clientOptions := options.Client().ApplyURI(configuration.Global.ConnectionString)
	return DataContext{ClientOptions: clientOptions}
}

func (ctx *DataContext) Collection(collection string) *mongo.Collection {
	if !ctx.connected {
		fmt.Println("Connecting")
		ctx.Client, _ = mongo.Connect(context.TODO(), ctx.ClientOptions)
		ctx.connected = true
	}
	return ctx.Client.Database(configuration.Global.Database).Collection(collection)
}

func (ctx *DataContext) Disconnect() error {

	if !ctx.connected {
		return nil
	}

	fmt.Println("Disconnecting")
	err := ctx.Client.Disconnect(context.TODO())

	if err != nil {
		return err
	}
	ctx.connected = false

	return nil
}
