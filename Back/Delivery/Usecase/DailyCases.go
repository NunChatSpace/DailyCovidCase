package Usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mainmodule/Database"
	"mainmodule/Delivery"
	"mainmodule/Domain"
	"mainmodule/Model"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type DailyCases struct {
	DB Model.MongoDBStruct
}

func (c *DailyCases) GetData(ctx *fiber.Ctx) Model.ResponseModel {
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
		return resp
	}

	resp = Model.ResponseModel{
		Status:  http.StatusOK,
		Message: "Getting data is successfully",
		Data:    content,
	}
	return resp
}

func NewDailyCases() Domain.DailyCasesInterface {
	apiToData := "https://covid19.th-stat.com/api/open/today"
	receivedData := &Model.DailyCases{}
	body, _ := Delivery.LoadData(apiToData)

	jsonErr := json.Unmarshal(body, receivedData)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return nil
	}

	dbStruct := Database.GetMongoDBStruct()
	context, _ := context.WithTimeout(context.Background(), 10*time.Second)

	dbStruct.DailyCasesCollection.DeleteMany(context, bson.D{{}})
	_, err := dbStruct.DailyCasesCollection.InsertOne(context, receivedData)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &DailyCases{
		DB: dbStruct,
	}
}
