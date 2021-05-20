package Domain

import (
	"mainmodule/Model"

	"github.com/gofiber/fiber/v2"
)

type ConvidCasesStatisticInterface interface {
	GetData(c *fiber.Ctx) Model.ResponseModel
}
