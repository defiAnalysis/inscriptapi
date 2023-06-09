package main

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"inscriptapi/models"
	"io/ioutil"
	"net/http"
	"strings"
)

type Data struct {
	Count    int    `json:"count"`
	More     bool   `json:"more"`
	Rows     []Row  `json:"rows"`
	Next     string `json:"next"`
	ScrollID string `json:"scrollId"`
}

type Row struct {
	Owner            Owner    `json:"owner"`
	ExpiryDate       string   `json:"expiryDate"`
	TokenID          string   `json:"tokenId"`
	Collections      []string `json:"collections"`
	RegistrationDate string   `json:"registrationDate"`
	Length           int      `json:"length"`
	LabelName        string   `json:"labelName"`
	CreateDate       string   `json:"createDate"`
	Charset          Charset  `json:"charset"`
	SegmentLength    int      `json:"segmentLength"`
	TopBid           TopBid   `json:"topBid"`
}

type Owner struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Charset struct {
	Emojis      int    `json:"emojis"`
	OnlyDigits  bool   `json:"onlyDigits"`
	OnlyLetters bool   `json:"onlyLetters"`
	Palindrome  bool   `json:"palindrome"`
	OnlyEmojis  bool   `json:"onlyEmojis"`
	Digits      int    `json:"digits"`
	Kinds       int    `json:"kinds"`
	Letters     int    `json:"letters"`
	Pattern     string `json:"pattern"`
}

type TopBid struct {
	Kind              string   `json:"kind"`
	Criteria          Criteria `json:"criteria"`
	Contract          string   `json:"contract"`
	Maker             string   `json:"maker"`
	ValidFrom         int      `json:"validFrom"`
	Source            string   `json:"source"`
	Nonce             string   `json:"nonce"`
	Price             float64  `json:"price"`
	TokenSetID        string   `json:"tokenSetId"`
	QuantityRemaining int      `json:"quantityRemaining"`
	ValidUntil        int      `json:"validUntil"`
	ID                string   `json:"id"`
	Value             float64  `json:"value"`
	Status            string   `json:"status"`
	UpdatedAt         int      `json:"updatedAt"`
}

type Criteria struct {
	Data DataCriteria `json:"data"`
	Kind string       `json:"kind"`
}

type DataCriteria struct {
	Token TokenCriteria `json:"token"`
}

type TokenCriteria struct {
	TokenID string `json:"tokenId"`
}

//
//type Params struct {
//	Id string `json:"id"`
//}
//
//type Filters struct {
//	Type   string      `json:"type"`
//	Params interface{} `json:"params"`
//}
//
//type RespData struct {
//	Filters  []Filters `json:"filters"`
//	Sort     string    `json:"sort"`
//	ScrollID string    `json:"scrollId"`
//	Next     string    `json:"next"`
//}

var (
	dataMap map[string]bool
)

func main() {
	beego.Info("===")
	dataMap = make(map[string]bool)

	session := orm.NewOrm()
	session.Begin()
	GetETHInfo("", "", session)

	ens := models.Ens{}

	beego.Info("ens:", ens)
	//url := "https://api.godid.io/api/items/query"
	//
	//data := `{"filters":[{"type":"activity_types","params":["sale"]},{"type":"collection","params":{"id":"b9Fby7ll5e"}},{"type":"status","params":{"status":"all"}}],"sort":"registrationDate_asc","scrollId":"","next":""}`
	//
	//s, err := HttpPostRaw(url, data)
	//if err != nil {
	//	fmt.Println("err:", err.Error())
	//	return
	//}
	//
	//var resp Data
	//if err = json.Unmarshal([]byte(s), &resp); err != nil {
	//	fmt.Println("error:", err)
	//	return
	//}
	//
	//beego.Info("====s=====:", resp)
	//
	//beego.Info("====len s=====:", len(resp.Rows))
	//
	//session := orm.NewOrm()
	//session.Begin()
	//
	//for i := 0; i < len(resp.Rows); i++ {
	//	record := resp.Rows[i]
	//
	//	ens := models.Ens{
	//		OwnerName:        record.Owner.Name,
	//		OwnerAddress:     record.Owner.Address,
	//		LabelName:        record.LabelName,
	//		RegistrationDate: record.RegistrationDate,
	//	}
	//
	//	if _, err := session.Insert(&ens); err != nil {
	//		beego.Error("err:", err)
	//	}
	//}

	session.Commit()
	return
}

func GetETHInfo(scrollId, next string, session orm.Ormer) string {
	url := "https://api.godid.io/api/items/query"

	data := `{"filters":[{"type":"activity_types","params":["sale"]},{"type":"collection","params":{"id":"etHpTUqIAk"}},{"type":"q","params":{"value":" "}},{"type":"status","params":{"status":"all"}}],"sort":"registrationDate_asc","scrollId":"` + scrollId + `","next":"` + next + `"}`

	if dataMap[next] {
		return ""
	} else {
		dataMap[next] = true
	}

	beego.Info("data:", data)
	s, err := HttpPostRaw(url, data)
	if err != nil {
		beego.Error("err:", err.Error())
		return ""
	}

	beego.Info("====s=====:", s)

	var resp Data
	if err := json.Unmarshal([]byte(s), &resp); err != nil {
		beego.Error("error:", err)
		return ""
	}

	if len(resp.Rows) == 0 {
		return ""
	}

	for i := 0; i < len(resp.Rows); i++ {
		record := resp.Rows[i]

		ens := models.Ens{
			OwnerName:        record.Owner.Name,
			OwnerAddress:     record.Owner.Address,
			LabelName:        record.LabelName + ".eth",
			RegistrationDate: record.RegistrationDate,
		}

		if _, err := session.Insert(&ens); err != nil {
			beego.Error("err:", err)
		}
	}

	return GetETHInfo(resp.ScrollID, resp.Next, session)
}

func HttpPostRaw(url, json_str string) (string, error) {
	payload := strings.NewReader(json_str)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("referer", "https://pancakeswap.finance/")
	//req.Header.Add("origin", "https://pancakeswap.finance/")
	res, err := http.DefaultClient.Do(req)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		return string(body), err
	}
	return "", err
}
