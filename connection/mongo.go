package connection

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var paymentCollection *mongo.Collection

func InitMongoDB() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}
	paymentCollection = client.Database("GO").Collection("Payment")
	return nil
}

func GetPaymentCollection() *mongo.Collection {
	return paymentCollection
}
