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

type CovidStat struct {
	DB Model.MongoDBStruct
}

func (c *CovidStat) GetData(ctx *fiber.Ctx) Model.ResponseModel {
	resp := Model.ResponseModel{}
	body := new(Model.SelectDateModel)
	filter := bson.D{}
	if ctx.Body() != nil {
		bodyParseError := ctx.BodyParser(body)
		if bodyParseError != nil {
			resp = Model.ResponseModel{
				Status:  http.StatusInternalServerError,
				Message: bodyParseError.Error(),
			}
			return resp
		}
		filter = c.makeFilter(body.FromDate, body.ToDate)
	}
	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := c.DB.CovidStatCollection.Find(context, filter)

	if err != nil {
		log.Fatal(err.Error())
		resp = Model.ResponseModel{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return resp
	}
	var dataContent []interface{}

	for cursor.Next(context) {
		var content Model.StatisticsEachTimeDay
		err := cursor.Decode(&content)
		// fmt.Println(content)
		if err != nil {
			resp = Model.ResponseModel{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
			return resp
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
	return resp
}

func (c *CovidStat) makeFilter(fromDate string, toDate string) bson.D {
	query := bson.D{}
	// bson.D{{"foo", "bar"}, {"hello", "world"}, {"pi", 3.14159}}
	fd, _ := time.Parse("01/02/2006", fromDate)
	td, _ := time.Parse("01/02/2006", toDate)
	if fromDate != "" && toDate != "" {
		query = bson.D{{"date", bson.D{{"$gte", fd}, {"$lte", td}}}}
	}
	return query
}

func NewCovidStat() Domain.ConvidCasesStatisticInterface {
	apiToData := "https://covid19.th-stat.com/api/open/timeline"
	receivedData := &Model.ConvidCasesStatistic{}
	body, _ := Delivery.LoadData(apiToData)
	jsonErr := json.Unmarshal(body, receivedData)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return nil
	}
	intf := make([]interface{}, len(receivedData.Data))

	for i, c := range receivedData.Data {
		date, _ := time.Parse("01/02/2006", c.Date)
		// fmt.Println(date)
		val := Model.StatisticsEachTimeDay{
			Date:            date,
			NewConfirmed:    c.NewConfirmed,
			NewRecovered:    c.NewRecovered,
			NewHospitalized: c.NewHospitalized,
			NewDeaths:       c.NewDeaths,
			Confirmed:       c.Confirmed,
			Recovered:       c.Recovered,
			Hospitalized:    c.Hospitalized,
			Deaths:          c.Deaths,
		}
		intf[i] = val
	}

	dbStruct := Database.GetMongoDBStruct()
	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	dbStruct.CovidStatCollection.DeleteMany(context, bson.D{{}})
	_, err := dbStruct.CovidStatCollection.InsertMany(context, intf)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &CovidStat{
		DB: dbStruct,
	}
}
