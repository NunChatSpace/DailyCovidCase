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

type CasesSum struct {
	DB Model.MongoDBStruct
}

func (c *CasesSum) GetData(ctx *fiber.Ctx) {
	resp := Model.ResponseModel{}
	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	body := new(Model.CasesSum)
	content := new(Model.CasesSum)
	jsonErr := json.Unmarshal([]byte(ctx.Body()), body)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	fmt.Println(ctx.Body())
	err := c.DB.CasesSumCollection.FindOne(context, bson.D{}).Decode(&content)

	if err != nil {
		log.Fatal(err.Error())
		resp = Model.ResponseModel{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		ctx.JSON(resp)
	}

	dataContent := c.makeFilter(content, body)
	resp = Model.ResponseModel{
		Status:  http.StatusOK,
		Message: "Getting data is successfully",
		Data:    dataContent,
	}
	ctx.JSON(resp)
}

func (c *CasesSum) makeFilter(bm *Model.CasesSum, km *Model.CasesSum) interface{} {
	result := make(map[string]interface{})

	if len(km.Province) > 0 {
		fmt.Println(km.Province)
		result["Province"] = c.getValue(bm.Province, km.Province)
	}
	if len(km.Nation) > 0 {
		fmt.Println(km.Nation)
		result["Nation"] = c.getValue(bm.Nation, km.Nation)
	}
	if len(km.Gender) > 0 {
		fmt.Println(km.Gender)
		result["Gender"] = c.getValue(bm.Gender, km.Gender)
	}

	return result
}

func (c *CasesSum) getValue(bm map[string]int, km map[string]int) interface{} {
	result := make(map[string]interface{})
	for k := range km {
		val := fmt.Sprintf("%d", bm[k])
		result[k] = val
		// result = append(result, val)
	}

	return result
}

func NewCasesSum() Domain.CasesSumInterface {
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

	return &CasesSum{
		DB: dbStruct,
	}
}
