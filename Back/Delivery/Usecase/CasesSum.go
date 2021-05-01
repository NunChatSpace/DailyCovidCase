package Usecase

import (
	"encoding/json"
	"mainmodule/Delivery"
	"mainmodule/Domain"
	"mainmodule/Model"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type casesSum struct {
	DB *gorm.DB
}

func (c *casesSum) GetData(ctx *fiber.Ctx) {
	ctx.SendString("CovidStatistics: " + "GetData")
}

func NewCasesSum(db *gorm.DB) Domain.CasesSumInterface {
	apiToData := "https://covid19.th-stat.com/api/open/cases/sum"
	receivedData := &Model.CasesSum{}
	body := Delivery.LoadData(apiToData)

	jsonErr := json.Unmarshal(body, receivedData)
	if jsonErr != nil {
		panic(jsonErr)
	}

	db.AutoMigrate(&Model.CasesSum{})
	tmpData := []Model.CasesSum{}
	db.Find(&tmpData)
	db.Where("1=1").Delete(&tmpData)

	db.Create(receivedData)
	return &casesSum{
		DB: db,
	}
}
