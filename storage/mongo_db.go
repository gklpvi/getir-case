package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"

	"getir-case/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBStorage struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoDBStorage(uri, dbName, collectionName string) (*MongoDBStorage, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to client: %v", err)
	}

	collection := client.Database(dbName).Collection(collectionName)

	return &MongoDBStorage{
		client:     client,
		collection: collection,
	}, nil
}

func (m *MongoDBStorage) Close() {
	err := m.client.Disconnect(context.Background())
	if err != nil {
		log.Printf("error while disconnecting from mongo: %v", err)
	}
}

func (m *MongoDBStorage) GetRecords(startDate, endDate time.Time, minCount, maxCount int) ([]model.DBRecord, error) {
	matchStage1 := bson.D{ //first match for range date time
		{Key: "$match", Value: bson.D{
			{Key: "createdAt", Value: bson.D{
				{Key: "$lte", Value: primitive.NewDateTimeFromTime(endDate)},
				{Key: "$gte", Value: primitive.NewDateTimeFromTime(startDate)}},
			},
		},
		},
	}
	projectStage := bson.D{ //sum counts as totalCount
		{Key: "$project", Value: bson.D{
			{Key: "totalCount", Value: bson.D{{Key: "$sum", Value: "$counts"}}},
			{Key: "_id", Value: 1},
			{Key: "key", Value: 1},
			{Key: "createdAt", Value: 1},
		},
		},
	}

	matchStage2 := bson.D{ //second match stage for total count
		{Key: "$match", Value: bson.D{
			{Key: "totalCount", Value: bson.D{
				{Key: "$lte", Value: maxCount},
				{Key: "$gte", Value: minCount}}},
		},
		},
	}
	pipeline := mongo.Pipeline{matchStage1, projectStage, matchStage2}

	cur, err := m.collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, fmt.Errorf("error while aggregating: %v", err)
	}

	defer cur.Close(context.Background())

	var records []model.DBRecord
	for cur.Next(context.Background()) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			return nil, fmt.Errorf("error while decoding: %v", err)
		}

		var record model.DBRecord
		record.Key = result.Map()["key"].(string)
		record.CreatedAt = result.Map()["createdAt"].(primitive.DateTime).Time().Format(time.RFC3339)
		record.TotalCount = result.Map()["totalCount"].(int)
		records = append(records, record)
	}

	return records, nil
}

// AddItem adds a new item to MongoDB
func (m *MongoDBStorage) AddItem(item *model.DBRecord) error {
	_, err := m.collection.InsertOne(context.Background(), item)
	return err
}
