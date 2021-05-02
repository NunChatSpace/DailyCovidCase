package Model

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBStruct struct {
	MongoDB              *mongo.Database
	CasesInfoCollection  *mongo.Collection
	CasesSumCollection   *mongo.Collection
	CovidStatCollection  *mongo.Collection
	DailyCasesCollection *mongo.Collection
	MongoDBClient        *mongo.Client
}
