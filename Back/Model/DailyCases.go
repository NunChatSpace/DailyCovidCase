package Model

import "gorm.io/gorm"

type DailyCases struct {
	*gorm.Model
	Confirmed       uint   `json:"Confirmed"`
	Recovered       uint   `json:"Recovered"`
	Hospitalized    uint   `json:"Hospitalized"`
	Deaths          uint   `json:"Deaths"`
	NewConfirmed    uint   `json:"NewConfirmed"`
	NewRecovered    uint   `json:"NewRecovered"`
	NewHospitalized uint   `json:"NewHospitalized"`
	NewDeaths       uint   `json:"NewDeaths"`
	UpdateDate      string `json:"UpdateDate"`
	Source          string `json:"Source"`
	DevBy           string `json:"DevBy"`
	SeverBy         string `json:"SeverBy"`
}
