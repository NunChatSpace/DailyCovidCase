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

type CovidStat struct {
	DB *gorm.DB
}

func (c *CovidStat) GetData(ctx *fiber.Ctx) {
	ctx.SendString("CovidStatistics: " + "GetData")
}

// func NewCovidStat(db *gorm.DB) Domain.ConvidCasesStatisticInterface {
// 	apiToData := "https://covid19.th-stat.com/api/open/timeline"
// 	receivedData := &Model.ConvidCasesStatistic{}
// 	body := Delivery.LoadData(apiToData)
// 	jsonErr := json.Unmarshal(body, receivedData)
// 	if jsonErr != nil {
// 		panic(jsonErr)
// 	}

// 	db.AutoMigrate(&Model.StatisticsEachDay{})
// 	tmpData := []Model.StatisticsEachDay{}
// 	db.Find(&tmpData)
// 	db.Where("1=1").Delete(&tmpData)

// 	db.Create(receivedData.Data)
// 	return &CovidStat{
// 		DB: db,
// 	}
// }
func NewCovidStat() Domain.CasesInfoInterface {
	apiToData := "https://covid19.th-stat.com/api/open/timeline"
	receivedData := &Model.ConvidCasesStatistic{}
	body := Delivery.LoadData(apiToData)

	jsonErr := json.Unmarshal(body, receivedData)
	if jsonErr != nil {
		panic(jsonErr)
	}
	intf := make([]interface{}, len(receivedData.Data))

	for i, c := range receivedData.Data {
		intf[i] = c
	}

	dbStruct := Database.GetMongoDBStruct()
	dbStruct.CovidStatCollection.DeleteMany(dbStruct.MongoDBContext, bson.D{{}})
	_, err := dbStruct.CovidStatCollection.InsertMany(dbStruct.MongoDBContext, intf)
	if err != nil {
		panic(err)
	}
	return &cases{
		DB: dbStruct,
	}
}
