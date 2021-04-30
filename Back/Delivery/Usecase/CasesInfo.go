package Usecase

import (
	"mainmodule/Domain"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type cases struct {
	DB *gorm.DB
}

func (c *cases) GetData(ctx *fiber.Ctx) {
	ctx.SendString("Cases Info: " + "GetData")
}

func NewCasesInfo(db *gorm.DB) Domain.CasesInfoInterface {
	return &cases{
		DB: db,
	}
}
