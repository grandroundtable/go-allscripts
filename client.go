package allscripts

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

const ehrtokenregex = "[A-Z0-9]{8}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{12}"

func checkParamValid(validvals []string, passed string) (err error) {
	for _, p := range validvals {
		if p == passed {
			return
		}
	}
	return errors.New(fmt.Sprintf("%s is not a valid value", passed))
}

// Client is a HTTP client of the Allscripts Professional API.
type Client struct {
	Userid      string
	Appname     string
	Appusername string
	Apppassword string
	Endpoint    string
	Token       string
}

// MagicJsonRequest is the JSON request body that accompanies /MagicJson API
// requests.
type MagicJsonRequest struct {
	Userid  string `json:"AppUserID"`
	Action  string `json:"Action"`
	Appname string `json:"Appname"`
	Patient string `json:"PatientID"`
	Token   string `json:"Token"`
	Param1  string `json:"Parameter1`
	Param2  string `json:"Parameter2`
	Param3  string `json:"Parameter3`
	Param4  string `json:"Parameter4`
	Param5  string `json:"Parameter5`
	Param6  string `json:"Parameter6`
	Data    string `json:"Data"`
}

// TokenRequest is the JSON request body that accompanies /GetToken API
// requests.
type TokenRequest struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

// NewClient creates a new Client and returns it.
func NewClient(userid string, appname string, appusername string,
	apppassword string, endpoint string) Client {
	return Client{userid, appname, appusername, apppassword, endpoint, ""}
}

func (c *Client) getToken() (token string, err error) {
	var (
		raw     []byte
		req     *http.Request
		res     *http.Response
		resRaw  []byte
		matched bool
	)

	tr := TokenRequest{
		c.Appusername,
		c.Apppassword,
	}

	raw, err = json.Marshal(tr)
	if err != nil {
		return "", errors.New("unable to get app token")
	}

	client := &http.Client{}
	req, err = http.NewRequest("POST", c.Endpoint+"/GetToken",
		bytes.NewReader(raw))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err = client.Do(req)
	if err != nil {
		return "", err
	}

	resRaw, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	res.Body.Close()

	matched, err = regexp.Match(ehrtokenregex, resRaw)
	if err != nil || !matched {
		return "", errors.New("unable to fetch GRT token from Allscripts")
	}

	token = string(resRaw)

	return
}

// ConstructReq constructs and returns a MagicJsonRequest struct that is used
// to make /MagicJson API requests.
func (c *Client) ConstructReq(action string, data map[string]string) (req MagicJsonRequest,
	err error) {
	if c.Token == "" {
		var token string
		token, err = c.getToken()
		if err != nil {
			return
		}
		c.Token = token
	}

	req = MagicJsonRequest{
		c.Userid,
		action,
		c.Appname,
		data["Patient"],
		c.Token,
		data["Param1"],
		data["Param2"],
		data["Param3"],
		data["Param4"],
		data["Param5"],
		data["Param6"],
		data["Data"],
	}

	return
}

// MakeRequest makes a request to the Allscripts Professional API.
func (c *Client) MakeRequest(reqbody MagicJsonRequest) (resp []byte,
	err error) {
	var (
		raw []byte
		req *http.Request
		res *http.Response
	)

	raw, err = json.Marshal(reqbody)
	if err != nil {
		return
	}

	client := &http.Client{}
	req, err = http.NewRequest("POST", c.Endpoint+"/MagicJson",
		bytes.NewReader(raw))
	if err != nil {
		return []byte(nil), err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err = client.Do(req)
	if err != nil {
		return []byte(nil), err
	}

	resp, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte(nil), err
	}
	res.Body.Close()

	return
}
