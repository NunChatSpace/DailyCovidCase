package Model

type ResponseModel struct {
	Status     int32       `json:"Status"`
	Message    string      `json:"Message"`
	DataLength int         `json:"DataLength"`
	Data       interface{} `json:"Data"`
}
