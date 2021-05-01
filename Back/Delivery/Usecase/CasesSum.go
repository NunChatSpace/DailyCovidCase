package Usecase

import (
	"encoding/json"
	"mainmodule/Database"
	"mainmodule/Delivery"
	"mainmodule/Domain"
	"mainmodule/Model"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"gorm.io/gorm"
)

type casesSum struct {
	DB *gorm.DB
}

func (c *casesSum) GetData(ctx *fiber.Ctx) {
	ctx.SendString("CovidStatistics: " + "GetData")
}

func NewCasesSum() Domain.CasesInfoInterface {
	apiToData := "https://covid19.th-stat.com/api/open/cases/sum"
	receivedData := &Model.CasesSum{}
	body := Delivery.LoadData(apiToData)

	jsonErr := json.Unmarshal(body, receivedData)
	if jsonErr != nil {
		panic(jsonErr)
	}

	dbStruct := Database.GetMongoDBStruct()
	dbStruct.CasesSumCollection.DeleteMany(dbStruct.MongoDBContext, bson.D{{}})
	_, err := dbStruct.CasesSumCollection.InsertOne(dbStruct.MongoDBContext, receivedData)
	if err != nil {
		panic(err)
	}

	return &cases{
		DB: dbStruct,
	}
}
