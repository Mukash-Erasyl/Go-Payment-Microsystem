package repository

import (
	"GoMicroservices/connection"
	"GoMicroservices/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentRepository struct {
	collection *mongo.Collection
}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{collection: connection.GetPaymentCollection()}
}

func (r *PaymentRepository) FindAll(ctx context.Context) ([]domain.Payment, error) {
	var payments []domain.Payment
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &payments); err != nil {
		return nil, err
	}
	return payments, nil
}
