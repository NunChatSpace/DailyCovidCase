package Usecase

import (
	"mainmodule/Domain"

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
	return &casesSum{
		DB: db,
	}
}
