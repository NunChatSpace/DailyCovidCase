package Model

type CasesSum struct {
	Province   map[string]int `json:"Province"`
	Nation     map[string]int `json:"Nation"`
	Gender     map[string]int `json:"Gender"`
	LastData   string         `json:"LastData"`
	UpdateDate string         `json:"UpdateDate"`
	Source     string         `json:"Source"`
	DevBy      string         `json:"DevBy"`
	SeverBy    string         `json:"SeverBy"`
}
