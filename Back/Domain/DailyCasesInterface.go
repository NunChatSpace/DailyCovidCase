package Domain

import (
	"mainmodule/Model"

	"github.com/gofiber/fiber/v2"
)

type DailyCasesInterface interface {
	GetData(c *fiber.Ctx) Model.ResponseModel
}
