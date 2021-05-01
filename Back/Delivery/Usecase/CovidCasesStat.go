package Usecase

import (
	"encoding/json"
	"mainmodule/Delivery"
	"mainmodule/Domain"
	"mainmodule/Model"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type CovidStat struct {
	DB *gorm.DB
}

func (c *CovidStat) GetData(ctx *fiber.Ctx) {
	ctx.SendString("CovidStatistics: " + "GetData")
}

func NewCovidStat(db *gorm.DB) Domain.ConvidCasesStatisticInterface {
	apiToData := "https://covid19.th-stat.com/api/open/timeline"
	receivedData := &Model.ConvidCasesStatistic{}
	body := Delivery.LoadData(apiToData)
	jsonErr := json.Unmarshal(body, receivedData)
	if jsonErr != nil {
		panic(jsonErr)
	}

	db.AutoMigrate(&Model.StatisticsEachDay{})
	tmpData := []Model.StatisticsEachDay{}
	db.Find(&tmpData)
	db.Where("1=1").Delete(&tmpData)

	db.Create(receivedData.Data)
	return &CovidStat{
		DB: db,
	}
}
