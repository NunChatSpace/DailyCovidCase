package Delivery

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"mainmodule/Model"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoadData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		return nil
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	return body
}

func ExtracData(col *mongo.Collection, client *mongo.Client) (resp Model.ResponseModel) {
	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := col.Find(context, bson.D{{}})
	fmt.Printf("%v \n", context)
	if err != nil {
		log.Fatal(err.Error())
		resp = Model.ResponseModel{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return resp
	}
	var dataContent []interface{}
	defer CloseCursor(context, cursor)

	for cursor.Next(context) {
		var content Model.Info
		err := cursor.Decode(&content)
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
		Message:    "Getting data is success fully",
		DataLength: len(dataContent),
		Data:       dataContent,
	}

	return resp
}

func CloseCursor(ctx context.Context, cursor *mongo.Cursor) {
	if cursor != nil {
		fmt.Println("cursor is closed.")
		cursor.Close(ctx)
	} else {
		fmt.Println("cursor isn't closed.")
	}
}
