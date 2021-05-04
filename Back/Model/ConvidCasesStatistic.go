package Model

import "time"

type ConvidCasesStatistic struct {
	UpdateDate string                    `json:"UpdateDate" gorm:"UpdateDate"`
	Source     string                    `json:"Source" gorm:"Source"`
	DevBy      string                    `json:"DevBy" gorm:"DevBy"`
	SeverBy    string                    `json:"SeverBy" gorm:"SeverBy"`
	Data       []StatisticsEachStringDay `json:"Data" gorm:"Data"`
}

type StatisticsEachStringDay struct {
	Date            string `json:"Date" gorm:"Date"`
	NewConfirmed    int32  `json:"NewConfirmed" gorm:"NewConfirmed"`
	NewRecovered    int32  `json:"NewRecovered" gorm:"NewRecovered"`
	NewHospitalized int32  `json:"NewHospitalized" gorm:"NewHospitalized"`
	NewDeaths       int32  `json:"NewDeaths" gorm:"NewDeaths"`
	Confirmed       int32  `json:"Confirmed" gorm:"Confirmed"`
	Recovered       int32  `json:"Recovered" gorm:"Recovered"`
	Hospitalized    int32  `json:"Hospitalized" gorm:"Hospitalized"`
	Deaths          int32  `json:"Deaths" gorm:"Deaths"`
}

type StatisticsEachTimeDay struct {
	Date            time.Time `json:"Date" gorm:"Date"`
	NewConfirmed    int32     `json:"NewConfirmed" gorm:"NewConfirmed"`
	NewRecovered    int32     `json:"NewRecovered" gorm:"NewRecovered"`
	NewHospitalized int32     `json:"NewHospitalized" gorm:"NewHospitalized"`
	NewDeaths       int32     `json:"NewDeaths" gorm:"NewDeaths"`
	Confirmed       int32     `json:"Confirmed" gorm:"Confirmed"`
	Recovered       int32     `json:"Recovered" gorm:"Recovered"`
	Hospitalized    int32     `json:"Hospitalized" gorm:"Hospitalized"`
	Deaths          int32     `json:"Deaths" gorm:"Deaths"`
}
