package Domain

import "github.com/gofiber/fiber"

type DailyCasesInterface interface {
	GetData(c *fiber.Ctx)
}
