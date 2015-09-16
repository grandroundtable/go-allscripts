package allscripts

// ClinicalSummary is a snapshot of a patient's clinical history.
type ClinicalSummary struct {
	Findings []Finding `json:"getclinicalsummaryinfo"`
}

// Finding is an individual finding from the patient's clinical summary.
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

// GetClinicalSummary returns a snapshot of a patient's clinical history.
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
	var reqbody MagicJsonRequest
	reqbody, err = c.constructReq("GetClinicalSummary", data)
	if err != nil {
		return findings, err
	}
	findings, err = c.MakeRequest(reqbody)
	return
}
