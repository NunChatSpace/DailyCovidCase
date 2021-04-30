package Usecase

import (
	"mainmodule/Domain"

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
	return &dailyCases{
		DB: db,
	}
}
