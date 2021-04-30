package Delivery

import (
	rd "mainmodule/RoutesDomain"

	"github.com/gofiber/fiber"
)

func SetupRoutes() {
	app := fiber.New()

	groupRoutes(app)

	app.Listen(8081)
}

func groupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/dailyCases", rd.DailyCasesInterface.GetData)
	v1.Get("/CasesStatistics", rd.ConvidCasesStatisticInterface.GetData)
	v1.Get("/CasesSum", rd.CasesSumInterface.GetData)
	v1.Get("/CasesInfo", rd.CasesInfoInterface.GetData)

}
