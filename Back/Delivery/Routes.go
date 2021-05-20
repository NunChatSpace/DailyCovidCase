package Delivery

import (
	intf "mainmodule/Domain"
	"mainmodule/Model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type RequestHandler struct {
	DC   intf.DailyCasesInterface
	CS   intf.ConvidCasesStatisticInterface
	CSum intf.CasesSumInterface
	CI   intf.CasesInfoInterface
}

var handler RequestHandler

func SetupRoutes(app *fiber.App,
	dc intf.DailyCasesInterface,
	cs intf.ConvidCasesStatisticInterface,
	csum intf.CasesSumInterface,
	ci intf.CasesInfoInterface) {

	handler = RequestHandler{
		DC:   dc,
		CS:   cs,
		CSum: csum,
		CI:   ci,
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/dailyCases", func(c *fiber.Ctx) error {
		resp := Model.ResponseModel{}
		if handler.DC == nil {
			resp = Model.ResponseModel{
				Status:  http.StatusInternalServerError,
				Message: "Database wasn't created",
			}
		} else {
			resp = handler.DC.GetData(c)
		}

		return c.JSON(resp)
	})

	v1.Get("/CasesStatistics", func(c *fiber.Ctx) error {
		resp := Model.ResponseModel{}
		if handler.DC == nil {
			resp = Model.ResponseModel{
				Status:  http.StatusInternalServerError,
				Message: "Database wasn't created",
			}
		} else {
			resp = handler.CS.GetData(c)
		}

		return c.JSON(resp)
	})

	v1.Get("/CasesSum", func(c *fiber.Ctx) error {
		resp := Model.ResponseModel{}
		if handler.DC == nil {
			resp = Model.ResponseModel{
				Status:  http.StatusInternalServerError,
				Message: "Database wasn't created",
			}
		} else {
			resp = handler.CSum.GetData(c)
		}

		return c.JSON(resp)
	})

	v1.Get("/CasesInfo", func(c *fiber.Ctx) error {
		resp := Model.ResponseModel{}
		if handler.DC == nil {
			resp = Model.ResponseModel{
				Status:  http.StatusInternalServerError,
				Message: "Database wasn't created",
			}
		} else {
			resp = handler.CI.GetData(c)
		}

		return c.JSON(resp)
	})
}
