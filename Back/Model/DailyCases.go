package Model

type DailyCases struct {
	Confirmed       uint   `json:"Confirmed" gorm:"Confirmed"`
	Recovered       uint   `json:"Recovered" gorm:"Recovered"`
	Hospitalized    uint   `json:"Hospitalized" gorm:"Hospitalized"`
	Deaths          uint   `json:"Deaths" gorm:"Deaths"`
	NewConfirmed    uint   `json:"NewConfirmed" gorm:"NewConfirmed"`
	NewRecovered    uint   `json:"NewRecovered" gorm:"NewRecovered"`
	NewHospitalized uint   `json:"NewHospitalized" gorm:"NewHospitalized"`
	NewDeaths       uint   `json:"NewDeaths" gorm:"NewDeaths"`
	UpdateDate      string `json:"UpdateDate" gorm:"UpdateDate"`
	Source          string `json:"Source" gorm:"Source"`
	DevBy           string `json:"DevBy" gorm:"DevBy"`
	SeverBy         string `json:"SeverBy" gorm:"SeverBy"`
}
