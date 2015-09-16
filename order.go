package allscripts

import "encoding/json"

// Orders is an array of orders returned by the GetOrders API call.
type Orders []Order

// Order includes details about the patient and order.
type Order struct {
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	PatientID     string `json:"patientid"`
	DOB           string `json:"dateofbirth"`
	Orderdate     string `json:"orderdate"`
	OrderID       string `json:"orderid"`
	OrderitemID   string `json:"OrderItemID"`
	Ordername     string `json:"ordername"`
	Orderstatus   string `json:"orderstatus"`
	Orderedbyname string `json:"orderedbyname"`
	Orderextid    string `json:"orderextid"`
	ReqID         string `json:"reqID"`
	RplID         string `js	on:"rplid"`
	Rplname       string `json:"rplname"`
	Rplcode       string `json:"rplcode"`
}

func (o *Orders) UnmarshalJSON(b []byte) error {
	type rawOrders Orders
	type wrapper struct {
		Orders []Order `json:"getordersinfo"`
	}

	var w []wrapper

	err := json.Unmarshal(b, &w)
	if err != nil {
		return err
	}
	*o = Orders(w[0].Orders)

	return nil
}

// GetOrders lists orders in the system for particular patients, date ranges,
// statuses, and order names.
//
//     sincedate: a string in the format of YYYY/MM/DD
//     ordermask: a wildcard matcher, for example, "%injection%"
//     status: Defaults to all status if not used.
//             Valid values are: "Undefined", "Ordered", "Preliminary", "Final",
//	           "Reviewed", "Pending", "Canceled"
func (c *Client) GetOrders(patientid string,
	sincedate string, ordermask string, status string) (findings []byte,
	err error) {
	validvals := []string{"", "Undefined", "Ordered", "Preliminary", "Final",
		"Reviewed", "Pending", "Canceled"}
	err = checkParamValid(validvals, status)
	if err != nil {
		return
	}

	if sincedate != "" {
		err = checkDateValid(sincedate)
		if err != nil {
			return
		}
	}

	var data map[string]string
	data = make(map[string]string)
	data["Patient"] = patientid
	data["Param1"] = sincedate
	data["Param2"] = ordermask
	data["Param3"] = status
	var reqbody magicJsonRequest
	reqbody, err = c.constructReq("GetOrders", data)
	if err != nil {
		return findings, err
	}
	findings, err = c.MakeRequest(reqbody)
	return
}
