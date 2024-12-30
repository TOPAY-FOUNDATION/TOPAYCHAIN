package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBStorage struct {
	Client     *mongo.Client
	Database   string
	Collection string
}

func NewDBStorage(uri, database, collection string) (*DBStorage, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	return &DBStorage{
		Client:     client,
		Database:   database,
		Collection: collection,
	}, nil
}

func (dbs *DBStorage) Save(key string, value interface{}) error {
	data, err := bson.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %v", err)
	}
	_, err = dbs.Client.Database(dbs.Database).Collection(dbs.Collection).UpdateOne(
		context.TODO(),
		bson.M{"_id": key},
		bson.M{"$set": bson.M{"data": data}},
		options.Update().SetUpsert(true),
	)
	return err
}

func (dbs *DBStorage) Load(key string, result interface{}) error {
	var doc struct {
		Data []byte `bson:"data"`
	}
	err := dbs.Client.Database(dbs.Database).Collection(dbs.Collection).FindOne(
		context.TODO(),
		bson.M{"_id": key},
	).Decode(&doc)
	if err != nil {
		return fmt.Errorf("failed to find key: %v", err)
	}
	return json.Unmarshal(doc.Data, result)
}

func (dbs *DBStorage) Delete(key string) error {
	_, err := dbs.Client.Database(dbs.Database).Collection(dbs.Collection).DeleteOne(
		context.TODO(),
		bson.M{"_id": key},
	)
	return err
}

func (dbs *DBStorage) ListKeys() ([]string, error) {
	cursor, err := dbs.Client.Database(dbs.Database).Collection(dbs.Collection).Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to list keys: %v", err)
	}
	defer cursor.Close(context.TODO())

	var keys []string
	for cursor.Next(context.TODO()) {
		var doc struct {
			ID string `bson:"_id"`
		}
		if err := cursor.Decode(&doc); err == nil {
			keys = append(keys, doc.ID)
		}
	}
	return keys, nil
}
