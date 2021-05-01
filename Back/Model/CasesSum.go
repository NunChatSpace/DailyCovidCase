package Model

type CasesSum struct {
	Province   map[string]int `json:"Province" gorm:"Province"`
	Nation     map[string]int `json:"Nation" gorm:"Nation"`
	Gender     map[string]int `json:"Gender" gorm:"Gender"`
	LastData   string         `json:"LastData" gorm:"LastData"`
	UpdateDate string         `json:"UpdateDate" gorm:"UpdateDate"`
	Source     string         `json:"Source" gorm:"Source"`
	DevBy      string         `json:"DevBy" gorm:"DevBy"`
	SeverBy    string         `json:"SeverBy" gorm:"SeverBy"`
}
