package Usecase

import (
	"encoding/json"
	"log"
	"mainmodule/Database"
	"mainmodule/Delivery"
	"mainmodule/Domain"
	"mainmodule/Model"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"gorm.io/gorm"
)

type dailyCases struct {
	DB *gorm.DB
}

func (c *dailyCases) GetData(ctx *fiber.Ctx) {
	ctx.SendString("CovidStatistics: " + "GetData")
}

// func NewDailyCases(db *gorm.DB) Domain.DailyCasesInterface {
// 	apiToData := "https://covid19.th-stat.com/api/open/today"
// 	receivedData := &Model.DailyCases{}
// 	body := Delivery.LoadData(apiToData)

// 	jsonErr := json.Unmarshal(body, receivedData)
// 	if jsonErr != nil {
// 		panic(jsonErr)
// 	}

// 	db.AutoMigrate(&Model.DailyCases{})
// 	tmpData := []Model.DailyCases{}
// 	db.Find(&tmpData)
// 	db.Where("1=1").Delete(&tmpData)
// 	db.Create(receivedData)

// 	return &dailyCases{
// 		DB: db,
// 	}
// }

func NewDailyCases() Domain.CasesInfoInterface {
	apiToData := "https://covid19.th-stat.com/api/open/today"
	receivedData := &Model.DailyCases{}
	body := Delivery.LoadData(apiToData)

	jsonErr := json.Unmarshal(body, receivedData)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	dbStruct := Database.GetMongoDBStruct()
	dbStruct.DailyCasesCollection.DeleteMany(dbStruct.MongoDBContext, bson.D{{}})
	_, err := dbStruct.DailyCasesCollection.InsertOne(dbStruct.MongoDBContext, receivedData)
	if err != nil {
		panic(err)
	}

	return &cases{
		DB: dbStruct,
	}
}
