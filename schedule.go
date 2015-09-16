package allscripts

import (
	"encoding/json"
	"errors"
)

const daterangeregex = `[0-9]{1,2}\/[0-9]{1,2}\/(19|20)[0-9]{2}\|[0-9]{1,2}\/[0-9]{1,2}\/(19|20)[0-9]{2}?`
const sincedateregex = `[0-9]{1,2}\/[0-9]{1,2}\/(19|20)[0-9]{2}`

// Appointments is an array of appointments return by the GetSchedule API call.
type Appointments []Appointment

// Appointment includes details about the patient, the time and date, status,
// organization and provider.
type Appointment struct {
	PatientID        string `json:"patientID"`
	Patientname      string `json:"Patient"`
	Patientfirstname string `json:"PatientFirstName"`
	Patientlastname  string `json:"PatientLastName"`
	Image            string `json:"base64image"`
	Datetime         string `json:"ApptTime"`
	Time             string `json:"AppTime2"`
	Status           string `json:"Status"`
	Statuscolor      string `json:"statuscolor"`
	Duration         string `json:"Duration"`
	Timein           string `json:"TimeIn"`
	Timestart        string `json:"TimeStart"`
	Timeout          string `json:"TimeOut"`
	OrgID            string `json:"OrganizationID"`
	OrgMRN           string `json:"organizationMRN"`
	Orgname          string `json:"OrgName"`
	ProviderID       string `json:"ProviderID"`
	Providername     string `json:"Provider"`
	Location         string `json:"LocationName"`
	SchedComment     string `json:"SchedComment"`
	Comment          string `json:"Comment"`
	Schedtext        string `json:"SchedText"`
	Type             string `json:"type"`
}

func (o *Appointments) UnmarshalJSON(b []byte) error {
	type rawAppointments Appointments
	type wrapper struct {
		Appointments []Appointment `json:"getscheduleinfo"`
	}

	var w []wrapper

	err := json.Unmarshal(b, &w)
	if err != nil {
		return err
	}
	*o = Appointments(w[0].Appointments)

	return nil
}

// GetSchedule lists appointments for the specified date or date range.
//
//     daterange: single date (for example, 3/24/2015), or you can enter a date
//                range separated by a pipe (3/24/2015|3/29/2015).
//                Date format is MM/DD/YYYY
//	   sincedate: a string in the format of MM/DD/YYYY
//     includepic: include a 45px × 45px JPG of the patient? Valid values are Y,
//                 or N. Defaults to N.
//	   user: the patient's MRN. For all patients, specify "All". Default is "All"
// 	   appttype: pipe-delimited list of desired appointment types.
func (c *Client) GetSchedule(daterange string, sincedate string,
	includepic string, user string, appttype string) (resp []byte,
	err error) {
	picvalidvals := []string{"", "Y", "N"}
	err = checkParamValid(picvalidvals, includepic)
	if err != nil {
		return
	}

	if daterange != "" {
		err = checkValid([]byte(daterange), daterangeregex)
		if err != nil {
			return
		}
	} else {
		err = errors.New("daterange parameter cannot be empty")
		return
	}

	if sincedate != "" {
		err = checkValid([]byte(sincedate), sincedateregex)
		if err != nil {
			return
		}
	}

	var data map[string]string
	data = make(map[string]string)
	data["Param1"] = daterange
	data["Param2"] = sincedate
	data["Param3"] = includepic
	data["Param4"] = user
	data["Param5"] = appttype
	var reqbody MagicJsonRequest
	reqbody, err = c.constructReq("GetSchedule", data)
	if err != nil {
		return resp, err
	}
	resp, err = c.MakeRequest(reqbody)
	return
}
