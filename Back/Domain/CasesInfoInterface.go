package RoutesDomain

import "github.com/gofiber/fiber"

type CasesInfoInterface interface {
	GetData(c *fiber.Ctx) (err error)
}
