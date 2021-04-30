package Usecase

import (
	"mainmodule/Domain"

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
	return &CovidStat{
		DB: db,
	}
}
