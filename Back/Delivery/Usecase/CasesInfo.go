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

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

type cases struct {
	DB Model.MongoDBStruct
}

func (c *cases) GetData(ctx *fiber.Ctx) {
	resp := Model.ResponseModel{}
	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := c.DB.CasesInfoCollection.Find(context, bson.D{{}})
	fmt.Printf("%v \n", context)
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
	context, _ := context.WithTimeout(context.Background(), 10*time.Second)

	dbStruct.CasesInfoCollection.DeleteMany(context, bson.D{{}})
	_, err := dbStruct.CasesInfoCollection.InsertMany(context, intf)
	if err != nil {
		panic(err)
	}

	return &cases{
		DB: dbStruct,
	}
}
