package Model

import "gorm.io/gorm"

type ConvidCasesStatistic struct {
	gorm.Model
	UpdateDate string            `json:"UpdateDate"`
	Source     string            `json:"Source"`
	DevBy      string            `json:"DevBy"`
	SeverBy    string            `json:"SeverBy"`
	Data       statisticsEachDay `json:"Data"`
}

type statisticsEachDay struct {
	Date            string `json:"Date"`
	NewConfirmed    string `json:"NewConfirmed"`
	NewRecovered    string `json:"NewRecovered"`
	NewHospitalized string `json:"NewHospitalized"`
	NewDeaths       string `json:"NewDeaths"`
	Confirmed       string `json:"Confirmed"`
	Recovered       string `json:"Recovered"`
	Hospitalized    string `json:"Hospitalized"`
	Deaths          string `json:"Deaths"`
}
