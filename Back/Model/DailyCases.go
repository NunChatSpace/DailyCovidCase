package Model

type DailyCases struct {
	Confirmed       int    `json:"Confirmed" gorm:"Confirmed"`
	Recovered       int    `json:"Recovered" gorm:"Recovered"`
	Hospitalized    int    `json:"Hospitalized" gorm:"Hospitalized"`
	Deaths          int    `json:"Deaths" gorm:"Deaths"`
	NewConfirmed    int    `json:"NewConfirmed" gorm:"NewConfirmed"`
	NewRecovered    int    `json:"NewRecovered" gorm:"NewRecovered"`
	NewHospitalized int    `json:"NewHospitalized" gorm:"NewHospitalized"`
	NewDeaths       int    `json:"NewDeaths" gorm:"NewDeaths"`
	UpdateDate      string `json:"UpdateDate" gorm:"UpdateDate"`
	Source          string `json:"Source" gorm:"Source"`
	DevBy           string `json:"DevBy" gorm:"DevBy"`
	SeverBy         string `json:"SeverBy" gorm:"SeverBy"`
}
