package main

import (
	"context"
	"log"
	"mainmodule/Database"
	"mainmodule/Delivery/Usecase"
	"time"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := fiber.New()
	// Database.SetupDatabase()
	// db := Database.GetDB()

	// ci := uc.NewCasesInfo(db)
	// csum := uc.NewCasesSum(db)
	// cs := uc.NewCovidStat(db)
	// dc := uc.NewDailyCases(db)

	// delivery.SetupRoutes(app, dc, cs, csum, ci)

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)
	Database.SetupMongoDB(client, ctx)
	Usecase.NewCasesInfo()
	// csum := uc.NewCasesSum(db)
	// cs := uc.NewCovidStat(db)
	// dc := uc.NewDailyCases(db)

	app.Listen(":8082")
}
