package Model

type CasesInfo struct {
	Data info `json:"Data"`
}

type info struct {
	ConfirmDate    string `json:"ConfirmDate"`
	No             string `json:"No"`
	Age            string `json:"Age"`
	Gender         string `json:"Gender"`
	GenderEn       string `json:"GenderEn"`
	Nation         string `json:"Nation"`
	NationEn       string `json:"NationEn"`
	Province       string `json:"Province"`
	ProvinceId     uint   `json:"ProvinceId"`
	District       string `json:"District"`
	ProvinceEn     string `json:"ProvinceEn"`
	Detail         string `json:"Detail"`
	StatQuarantine uint   `json:"StatQuarantine"`
}
