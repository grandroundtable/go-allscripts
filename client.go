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
const dateregex = `^(19|20)[0-9]{1,2}\/[0-9]{1,2}\/[0-9]{1,2}`

func checkParamValid(validvals []string, passed string) (err error) {
	for _, p := range validvals {
		if p == passed {
			return
		}
	}
	return errors.New(fmt.Sprintf("%s is not a valid value", passed))
}

func checkDateValid(date string) (err error) {
	var matched bool
	matched, err = regexp.Match(dateregex, []byte(date))
	if err != nil || !matched {
		return errors.New("not a valid date")
	}

	return
}

func checkValid(val []byte, regex string) (err error) {
	var matched bool
	matched, err = regexp.Match(regex, val)
	if err != nil || !matched {
		return errors.New("not a valid value")
	}

	return
}

// Client is a HTTP client of the Allscripts Professional API.
type Client struct {
	Userid      string
	Appname     string
	Appusername string
	Apppassword string
	Endpoint    string
	Token       string
	HTTPClient  *http.Client
}

type magicJsonRequest struct {
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

type tokenRequest struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

// NewClient creates a new Client and returns it.
func NewClient(userid string, appname string, appusername string,
	apppassword string, endpoint string) Client {
	return Client{userid, appname, appusername, apppassword,
		endpoint, "", new(http.Client)}
}

// WithHTTPClient sets a Client's HTTP client for making HTTP requests, returning
// a Client pointer for chaining.
func (c *Client) WithHTTPClient(hc *http.Client) *Client {
	c.HTTPClient = hc
	return c
}

// GetToken uses the Allscripts' /GetToken resource to get a token that may be
// used in MagicJson calls.
func (c *Client) GetToken() (token string, err error) {
	var (
		raw     []byte
		req     *http.Request
		res     *http.Response
		resRaw  []byte
		matched bool
	)

	tr := tokenRequest{
		c.Appusername,
		c.Apppassword,
	}

	raw, err = json.Marshal(tr)
	if err != nil {
		return "", errors.New("unable to get app token")
	}

	req, err = http.NewRequest("POST", c.Endpoint+"/GetToken",
		bytes.NewReader(raw))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err = c.HTTPClient.Do(req)
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

func (c *Client) constructReq(action string, data map[string]string) (req magicJsonRequest,
	err error) {
	if c.Token == "" {
		var token string
		token, err = c.GetToken()
		if err != nil {
			return
		}
		c.Token = token
	}

	req = magicJsonRequest{
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

func (c *Client) makeRequest(reqbody magicJsonRequest) (resp []byte,
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

	req, err = http.NewRequest("POST", c.Endpoint+"/MagicJson",
		bytes.NewReader(raw))
	if err != nil {
		return []byte(nil), err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err = c.HTTPClient.Do(req)
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
