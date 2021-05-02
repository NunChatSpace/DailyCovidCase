package main

import (
	"context"
	"fmt"
	"log"
	"mainmodule/Database"
	"mainmodule/Delivery"
	"mainmodule/Delivery/Usecase"
	"time"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := fiber.New()

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(ctx)
	Database.SetupMongoDB(client)
	ci := Usecase.NewCasesInfo()
	csum := Usecase.NewCasesSum()
	cs := Usecase.NewCovidStat()
	dc := Usecase.NewDailyCases()
	Delivery.SetupRoutes(app, dc, cs, csum, ci)

	t := time.Now()
	fmt.Printf("Server start at: %s \n", t.String())
	log.Fatal(app.Listen(":9090"))
}
