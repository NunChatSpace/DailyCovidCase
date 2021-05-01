package Usecase

import (
	"encoding/json"
	"mainmodule/Delivery"
	"mainmodule/Domain"
	"mainmodule/Model"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type dailyCases struct {
	DB *gorm.DB
}

func (c *dailyCases) GetData(ctx *fiber.Ctx) {
	ctx.SendString("CovidStatistics: " + "GetData")
}

func NewDailyCases(db *gorm.DB) Domain.DailyCasesInterface {
	apiToData := "https://covid19.th-stat.com/api/open/today"
	receivedData := &Model.DailyCases{}
	body := Delivery.LoadData(apiToData)

	jsonErr := json.Unmarshal(body, receivedData)
	if jsonErr != nil {
		panic(jsonErr)
	}

	db.AutoMigrate(&Model.DailyCases{})
	tmpData := []Model.DailyCases{}
	db.Find(&tmpData)
	db.Where("1=1").Delete(&tmpData)
	db.Create(receivedData)

	return &dailyCases{
		DB: db,
	}
}
