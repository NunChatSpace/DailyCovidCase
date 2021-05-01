package Model

type CasesInfo struct {
	Data []Info `json:"Data"`
}

type Info struct {
	ConfirmDate    string  `json:"ConfirmDate" gorm:"ConfirmDate"`
	No             string  `json:"No" gorm:"No"`
	Age            float32 `json:"Age" gorm:"Age"`
	Gender         string  `json:"Gender" gorm:"Gender"`
	GenderEn       string  `json:"GenderEn" gorm:"GenderEn"`
	Nation         string  `json:"Nation" gorm:"Nation"`
	NationEn       string  `json:"NationEn" gorm:"NationEn"`
	Province       string  `json:"Province" gorm:"Province"`
	ProvinceId     uint    `json:"ProvinceId" gorm:"ProvinceId"`
	District       string  `json:"District" gorm:"District"`
	ProvinceEn     string  `json:"ProvinceEn" gorm:"ProvinceEn"`
	Detail         string  `json:"Detail" gorm:"Detail"`
	StatQuarantine uint    `json:"StatQuarantine" gorm:"StatQuarantine"`
}
