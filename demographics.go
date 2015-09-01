package allscripts

import (
	"encoding/json"
)

// Patient is the details of a patient.
type Patient struct {
	Firstname        string `json:"Firstname"`
	Middlename       string `json:"middlename"`
	Lastname         string `json:"LastName"`
	Suffix           string `json:"suffix"`
	Mrn              string `json:"mrn"`
	SSN              string `json:"ssn"`
	Gender           string `json:"gender"`
	Race             string `json:"Race"`
	Agedec           string `json:"AgeDec"`
	Age              string `json:"age"`
	DOB              string `json:"dateofbirth"`
	Maritalstatus    string `json:"MaritalStatus"`
	Image            string `json:"base64image"`
	Addressline1     string `json:"Addressline1"`
	Addressline2     string `json:"AddressLine2"`
	City             string `json:"City"`
	State            string `json:"State"`
	Zip              string `json:"ZipCode"`
	Phone            string `json:"PhoneNumber"`
	Homephone        string `json:"HomePhone"`
	Workphone        string `json:"WorkPhone"`
	Cellphone        string `json:"cellphone"`
	Language         string `json:"Language"`
	Email            string `json:"Email"`
	Homeemail        string `json:"HomeEmail"`
	Physusername     string `json:"PhysUserName"`
	Physfirstname    string `json:"PhysFirstName"`
	Physlastname     string `json:"PhysLastName"`
	Physphone        string `json:"PhysPhone"`
	Location         string `json:"PatientLocation"`
	Primaryinsurance string `json:"PrimaryInsurance"`
	DEM              string `json:"DEM_EXTERNALID"`	
}

func (p *Patient) UnmarshalJSON(b []byte) error {
	type rawPatient Patient
	type wrapper struct {
		Patients []rawPatient `json:"getpatientinfo"`
	}

	var w []wrapper

	err := json.Unmarshal(b, &w)
	if err != nil {
		return err
	}
	*p = Patient(w[0].Patients[0])

	return nil
}

// GetPatient returns information for the specified Patient ID.
func (c *Client) GetPatient(patientid string) (patient []byte, err error) {
	var data map[string]string
	data = make(map[string]string)
	data["Patient"] = patientid
	var reqbody MagicJsonRequest
	reqbody, err = c.ConstructReq("GetPatient", data)
	if err != nil {
		return patient, err
	}

	patient, err = c.MakeRequest(reqbody)
	return
}
