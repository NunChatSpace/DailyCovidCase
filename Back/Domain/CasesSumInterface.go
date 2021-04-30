package RoutesDomain

import "github.com/gofiber/fiber"

type CasesSumInterface interface {
	GetData(c *fiber.Ctx) (err error)
}
