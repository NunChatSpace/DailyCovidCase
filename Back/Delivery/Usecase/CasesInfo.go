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

type cases struct {
	DB Model.MongoDBStruct
}

func (c *cases) GetData(ctx *fiber.Ctx) {
	resp := Model.ResponseModel{}
	body := new(Model.Info)
	bodyParseError := ctx.BodyParser(body)
	if bodyParseError != nil {
		resp = Model.ResponseModel{
			Status:  http.StatusInternalServerError,
			Message: bodyParseError.Error(),
		}
		ctx.JSON(resp)
	}
	filter := c.makeFilter(body)
	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := c.DB.CasesInfoCollection.Find(context, filter)
	// fmt.Printf("%v \n", context)
	if err != nil {
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

func (c *cases) makeFilter(m *Model.Info) bson.M {
	query := bson.M{}
	// bson.D{{"foo", "bar"}, {"hello", "world"}, {"pi", 3.14159}}
	if m.ConfirmDate != "" {
		query["confirmdate"] = m.ConfirmDate
	}
	if m.No != "" {
		query["no"] = m.No
	}
	if m.Age != 0 {
		query["age"] = m.Age
	}
	if m.Gender != "" {
		query["gender"] = m.Gender
	}
	if m.GenderEn != "" {
		query["genderen"] = m.GenderEn
	}
	if m.Nation != "" {
		query["nation"] = m.Nation
	}
	if m.NationEn != "" {
		query["nationen"] = m.NationEn
	}
	if m.Province != "" {
		query["province"] = m.Province
	}
	if m.ProvinceId != 0 {
		query["provinceid"] = m.ProvinceId
	}
	if m.District != "" {
		query["district"] = m.District
	}
	if m.ProvinceEn != "" {
		query["provinceen"] = m.ProvinceEn
	}
	if m.Detail != "" {
		query["detail"] = m.Detail
	}
	if m.StatQuarantine != 0 {
		query["statquarantine"] = m.StatQuarantine
	}
	// fmt.Println(query)
	return query
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
