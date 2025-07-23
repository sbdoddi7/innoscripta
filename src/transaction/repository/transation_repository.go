package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/sbdoddi7/innoscripta/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type transactionRepository struct {
	collection *mongo.Collection
	db         *sql.DB
}

func NewTransactionRepository(client *mongo.Client, dbName, collectionName string, db *sql.DB) *transactionRepository {
	return &transactionRepository{
		collection: client.Database(dbName).Collection(collectionName),
		db:         db,
	}
}

func (r *transactionRepository) WriteLog(log model.TransactionLog) error {
	log.Timestamp = time.Now()
	_, err := r.collection.InsertOne(context.TODO(), log)

	return err
}

func (r *transactionRepository) UpdateBalance(accountNumber int64, delta float64) error {
	query := `
	    UPDATE accounts
	    SET balance = balance + $1
	    WHERE account_number = $2
	`
	_, err := r.db.Exec(query, delta, accountNumber)
	return err
}

func (r *transactionRepository) GetTransactions(accountNumber int64, limit, offset int64) ([]model.TransactionLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"account_number": accountNumber}
	opts := options.Find().SetLimit(limit).SetSkip(offset).SetSort(bson.D{{Key: "timestamp", Value: -1}})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var logs []model.TransactionLog
	if err = cursor.All(ctx, &logs); err != nil {
		return nil, err
	}
	return logs, nil
}
