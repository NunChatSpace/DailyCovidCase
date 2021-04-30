package Delivery

import (
	intf "mainmodule/Domain"

	"github.com/gofiber/fiber"
)

type RequestHandler struct {
	DC   intf.DailyCasesInterface
	CS   intf.ConvidCasesStatisticInterface
	CSum intf.CasesSumInterface
	CI   intf.CasesInfoInterface
}

var handler RequestHandler

func SetupRoutes(app *fiber.App, dc intf.DailyCasesInterface, cs intf.ConvidCasesStatisticInterface, csum intf.CasesSumInterface, ci intf.CasesInfoInterface) {

	handler := &RequestHandler{
		DC:   dc,
		CS:   cs,
		CSum: csum,
		CI:   ci,
	}

	app.Get("/", func(c *fiber.Ctx) {
		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/dailyCases", handler.DC.GetData)
	v1.Get("/CasesStatistics", handler.CS.GetData)
	v1.Get("/CasesSum", handler.CSum.GetData)
	v1.Get("/CasesInfo", handler.CI.GetData)
}
