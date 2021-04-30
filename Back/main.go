package main

import (
	"mainmodule/Database"
	delivery "mainmodule/Delivery"
	uc "mainmodule/Delivery/Usecase"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	db := Database.GetDB()
	ci := uc.NewCasesInfo(db)
	csum := uc.NewCasesSum(db)
	cs := uc.NewCovidStat(db)
	dc := uc.NewDailyCases(db)

	delivery.SetupRoutes(app, dc, cs, csum, ci)

	app.Listen(":8081")
}
