package Database

import (
	"context"
	"mainmodule/Model"

	"go.mongodb.org/mongo-driver/mongo"
)

var mongoDB Model.MongoDBStruct

func GetMongoDBStruct() Model.MongoDBStruct {
	return mongoDB
}

func SetupMongoDB(client *mongo.Client, ctx context.Context) {
	dbTemp := client.Database("mainDatabase")
	mongoDB = Model.MongoDBStruct{
		MongoDB:              dbTemp,
		CasesInfoCollection:  dbTemp.Collection("CasesInfo"),
		CasesSumCollection:   dbTemp.Collection("CasesSum"),
		CovidStatCollection:  dbTemp.Collection("CovidStat"),
		DailyCasesCollection: dbTemp.Collection("DailyCases"),
		MongoDBContext:       ctx,
	}
}
