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

type casesSum struct {
	DB Model.MongoDBStruct
}

func (c *casesSum) GetData(ctx *fiber.Ctx) {
	resp := Model.ResponseModel{}
	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := c.DB.CasesSumCollection.Find(context, bson.D{{}})

	if err != nil {
		log.Fatal(err.Error())
		resp = Model.ResponseModel{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		ctx.JSON(resp)
	}
	var dataContent []interface{}

	for cursor.Next(context) {
		var content Model.Info
		err := cursor.Decode(&content)
		if err != nil {
			resp = Model.ResponseModel{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
			ctx.JSON(resp)
		}
		dataContent = append(dataContent, content)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	resp = Model.ResponseModel{
		Status:     http.StatusOK,
		Message:    "Getting data is successfully",
		DataLength: len(dataContent),
		Data:       dataContent,
	}
	ctx.JSON(resp)
}

func NewCasesSum() Domain.CasesInfoInterface {
	apiToData := "https://covid19.th-stat.com/api/open/cases/sum"
	receivedData := &Model.CasesSum{}
	body := Delivery.LoadData(apiToData)

	jsonErr := json.Unmarshal(body, receivedData)
	if jsonErr != nil {
		panic(jsonErr)
	}

	dbStruct := Database.GetMongoDBStruct()
	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	dbStruct.CasesSumCollection.DeleteMany(context, bson.D{{}})
	_, err := dbStruct.CasesSumCollection.InsertOne(context, receivedData)
	if err != nil {
		panic(err)
	}

	return &cases{
		DB: dbStruct,
	}
}
