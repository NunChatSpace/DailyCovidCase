package Domain

import "github.com/gofiber/fiber"

type CasesInfoInterface interface {
	GetData(c *fiber.Ctx)
}
