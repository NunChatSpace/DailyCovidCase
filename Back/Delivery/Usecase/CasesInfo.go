package Usecase

import (
	"encoding/json"
	"fmt"
	"mainmodule/Database"
	"mainmodule/Delivery"
	"mainmodule/Domain"
	"mainmodule/Model"

	"github.com/gofiber/fiber"
)

type cases struct {
	DB Model.MongoDBStruct
}

func (c *cases) GetData(ctx *fiber.Ctx) {
	ctx.SendString("Cases Info: " + "GetData")
}

// func NewCasesInfo(db *gorm.DB) Domain.CasesInfoInterface {
// 	apiToData := "https://covid19.th-stat.com/api/open/cases"
// 	receivedData := &Model.CasesInfo{}
// 	body := Delivery.LoadData(apiToData)

// 	jsonErr := json.Unmarshal(body, receivedData)
// 	if jsonErr != nil {
// 		panic(jsonErr)
// 	}

// 	db.AutoMigrate(&Model.Info{})
// 	tmpData := []Model.Info{}
// 	db.Find(&tmpData)
// 	db.Where("1=1").Delete(&tmpData)

// 	db.Create(receivedData.Data)
// 	return &cases{
// 		DB: db,
// 	}
// }

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

	fmt.Printf("intf length: %d\n", len(intf))
	dbStruct := Database.GetMongoDBStruct()
	insertResult, err := dbStruct.CasesInfoCollection.InsertMany(dbStruct.MongoDBContext, intf)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(insertResult)
	}
	return &cases{
		DB: dbStruct,
	}
}
