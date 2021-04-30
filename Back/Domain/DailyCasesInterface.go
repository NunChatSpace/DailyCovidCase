package RoutesDomain

import "github.com/gofiber/fiber"

type DailyCasesInterface interface {
	GetData(c *fiber.Ctx) (err error)
}
