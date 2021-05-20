package Domain

import (
	"mainmodule/Model"

	"github.com/gofiber/fiber/v2"
)

type CasesSumInterface interface {
	GetData(c *fiber.Ctx) Model.ResponseModel
}
