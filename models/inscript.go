package models

type Inscript struct {
	Id          int64  `orm:"pk;auto;description(主键id)"  form:"id" json:"id" `
	Type        string `json:"type"`
	Sname       string `json:"sname"`
	Oname       string `json:"oname"`
	FullName    string `json:"fullName"`
	Suffix      string `json:"suffix"`
	DomainName  string `json:"domainName"`
	Content     string `json:"content"`
	Categories  string `json:"categories"`
	Width       int    `json:"width"`
	Owner       string `json:"owner"`
	Inscription string `json:"inscription"`
	Num         int64  `json:"num"`
	OutputValue int    `json:"outputValue"`
	CreateTime  string `json:"createTime"`
	Status      int    `json:"status"`
}

func (a *Inscript) TableName() string {
	return InscriptTBName()
}
