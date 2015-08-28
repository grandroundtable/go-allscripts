package allscripts

import (
//"fmt"
)

type ClinicalSummary struct {
	Findings []Finding `json:"getclinicalsummaryinfo"`
}

type Finding struct {
	Detail      string `json:"detail"`
	Code        string `json:"code"`
	Transid     string `json:"transid"`
	Status      string `json:"status"`
	XMLDetail   string `json:"xmldetail"`
	Entrycode   string `json:"entrycode"`
	Description string `json:"description"`
	Section     string `json:"section"`
	Displaydate string `json:"displaydate"`
}

func (c *Client) GetClinicalSummary(patientid string,
	section string) (findings []byte, err error) {
	validvals := []string{"allergies", "history", "immunizations",
		"medications", "problems", "vitals", "results", "list"}
	err = checkParamValid(validvals, section)
	if err != nil {
		return
	}

	var data map[string]string
	data = make(map[string]string)
	data["Patient"] = patientid
	data["Param1"] = section
	reqbody, _ := c.ConstructReq("GetClinicalSummary", data)
	findings, err = c.MakeRequest(reqbody)
	return
}
