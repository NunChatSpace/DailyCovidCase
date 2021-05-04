package Usecase

import (
	"context"
	"encoding/json"
	"log"
	"mainmodule/Database"
	"mainmodule/Delivery"
	"mainmodule/Domain"
	"mainmodule/Model"
	"net/http"
	"time"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

type DailyCases struct {
	DB Model.MongoDBStruct
}

func (c *DailyCases) GetData(ctx *fiber.Ctx) {
	resp := Model.ResponseModel{}
	content := new(Model.DailyCases)

	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := c.DB.DailyCasesCollection.FindOne(context, bson.D{{}}).Decode(&content)
	if err != nil {
		log.Fatal(err.Error())
		resp = Model.ResponseModel{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		ctx.JSON(resp)
	}

	resp = Model.ResponseModel{
		Status:  http.StatusOK,
		Message: "Getting data is successfully",
		Data:    content,
	}
	ctx.JSON(resp)
}

func NewDailyCases() Domain.DailyCasesInterface {
	apiToData := "https://covid19.th-stat.com/api/open/today"
	receivedData := &Model.DailyCases{}
	body := Delivery.LoadData(apiToData)

	jsonErr := json.Unmarshal(body, receivedData)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	dbStruct := Database.GetMongoDBStruct()
	context, _ := context.WithTimeout(context.Background(), 10*time.Second)

	dbStruct.DailyCasesCollection.DeleteMany(context, bson.D{{}})
	_, err := dbStruct.DailyCasesCollection.InsertOne(context, receivedData)
	if err != nil {
		panic(err)
	}

	return &DailyCases{
		DB: dbStruct,
	}
}
