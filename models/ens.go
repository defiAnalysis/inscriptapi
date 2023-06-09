package models

type Ens struct {
	Id               int64  `orm:"pk;auto;description(主键id)"  form:"id" json:"id" `
	OwnerAddress     string `json:"address"`
	OwnerName        string `json:"name"`
	LabelName        string `json:"labelName"`
	RegistrationDate string `json:"registrationDate"`
}

func (a *Ens) TableName() string {
	return EnsTBName()
}
