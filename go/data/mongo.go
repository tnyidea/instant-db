package data

import "go.mongodb.org/mongo-driver/mongo"

const AddressCollection = "address"

type MongoConnection struct {
	Client   *mongo.Client
	Database *mongo.Database

	AddressCollection *mongo.Collection
}
