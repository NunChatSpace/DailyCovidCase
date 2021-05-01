package Model

type ConvidCasesStatistic struct {
	UpdateDate string              `json:"UpdateDate" gorm:"UpdateDate"`
	Source     string              `json:"Source" gorm:"Source"`
	DevBy      string              `json:"DevBy" gorm:"DevBy"`
	SeverBy    string              `json:"SeverBy" gorm:"SeverBy"`
	Data       []StatisticsEachDay `json:"Data" gorm:"Data"`
}

type StatisticsEachDay struct {
	Date            string `json:"Date" gorm:"Date"`
	NewConfirmed    int16  `json:"NewConfirmed" gorm:"NewConfirmed"`
	NewRecovered    int16  `json:"NewRecovered" gorm:"NewRecovered"`
	NewHospitalized int16  `json:"NewHospitalized" gorm:"NewHospitalized"`
	NewDeaths       int16  `json:"NewDeaths" gorm:"NewDeaths"`
	Confirmed       int16  `json:"Confirmed" gorm:"Confirmed"`
	Recovered       int16  `json:"Recovered" gorm:"Recovered"`
	Hospitalized    int16  `json:"Hospitalized" gorm:"Hospitalized"`
	Deaths          int16  `json:"Deaths" gorm:"Deaths"`
}
