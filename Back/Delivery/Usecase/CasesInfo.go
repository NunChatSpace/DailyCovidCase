package Usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"mainmodule/Database"
	"mainmodule/Delivery"
	"mainmodule/Domain"
	"mainmodule/Model"
	"net/http"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

type cases struct {
	DB Model.MongoDBStruct
}

func (c *cases) GetData(ctx *fiber.Ctx) {
	cursor, err := c.DB.CasesInfoCollection.Find(c.DB.MongoDBContext, bson.D{{}})
	resp := Model.ResponseModel{}
	if err != nil {
		resp = Model.ResponseModel{
			Status:  http.StatusInternalServerError,
			Message: "Data not found",
		}

		ctx.JSON(resp)
	}
	var dataContent []interface{}
	for cursor.Next(c.DB.MongoDBContext) {
		var content Model.Info
		err := cursor.Decode(&content)
		if err != nil {
			resp = Model.ResponseModel{
				Status:  http.StatusInternalServerError,
				Message: "Data not found",
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
		Message:    "Getting data is success fully",
		DataLength: len(dataContent),
		Data:       dataContent,
	}

	ctx.JSON(resp)
}

func NewCasesInfo() Domain.CasesInfoInterface {
	apiToData := "https://covid19.th-stat.com/api/open/cases"
	receivedData := &Model.CasesInfo{}
	body := Delivery.LoadData(apiToData)
	jsonErr := json.Unmarshal(body, receivedData)
	if jsonErr != nil {
		panic(jsonErr)
	}

	intf := make([]interface{}, len(receivedData.Data))

	for i, c := range receivedData.Data {
		intf[i] = c
	}
	dbStruct := Database.GetMongoDBStruct()

	dbStruct.CasesInfoCollection.DeleteMany(dbStruct.MongoDBContext, bson.D{{}})
	_, err := dbStruct.CasesInfoCollection.InsertMany(dbStruct.MongoDBContext, intf)
	if err != nil {
		panic(err)
	}

	return &cases{
		DB: dbStruct,
	}
}
