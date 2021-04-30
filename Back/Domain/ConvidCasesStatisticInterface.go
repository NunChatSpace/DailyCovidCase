package Domain

import "github.com/gofiber/fiber"

type ConvidCasesStatisticInterface interface {
	GetData(c *fiber.Ctx)
}
