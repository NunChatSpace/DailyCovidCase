package Domain

import "github.com/gofiber/fiber"

type CasesSumInterface interface {
	GetData(c *fiber.Ctx)
}
