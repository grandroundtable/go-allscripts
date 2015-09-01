package mock

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/onsi/gomega/ghttp"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"runtime"
)

var (
	validusername string
	validpassword string
	mockHandlers  map[string]http.HandlerFunc
)

const (
	token = "5B6FBE1C-6B9C-43AB-8A60-AD8E45AE813D"
)

// LoadSampleJson loads a file of JSON data that should be sent back to
// successful API requests.
func LoadSampleJson(f string) (out []byte, err error) {
	_, filename, _, _ := runtime.Caller(1)
	file := path.Join(path.Dir(filename), "sample_data", f)
	out, err = ioutil.ReadFile(file)
	return
}

// AddMockHandler adds a mock HTTP handler to the map.
func AddMockHandler(handler http.HandlerFunc, action string) {
	if mockHandlers == nil {
		mockHandlers = make(map[string]http.HandlerFunc)
	}
	mockHandlers[action] = handler
}

func getMockHandler(action string) (handler http.HandlerFunc, err error) {
	handler = mockHandlers[action]
	if handler == nil {
		err = errors.New("No handler found for that action")
		return nil, err
	}

	return
}

// TokenLogin is the JSON body structure sent for /GetToken API requests.
type TokenLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// MagicJsonBody is the JSON body structure sent from /MagicJson API requests.
type MagicJsonBody struct {
	Appuserid string `json:"AppUserID"`
	Action    string `json:"Action"`
	Cppname   string `json:"Appname"`
	Patientid string `json:"PatientID"`
	Token     string `json:"Token"`
	Param1    string `json:"Parameter1"`
	Param2    string `json:"Parameter2"`
	Param3    string `json:"Parameter3"`
	Param4    string `json:"Parameter4"`
	Param5    string `json:"Parameter5"`
	Param6    string `json:"Parameter6"`
	Data      string `json:"Data"`
}

// MockHandler is a HTTP handler that responds to requests to the mock server
// at a specific path
type MockHandler struct {
	Method      string
	Path        string
	HandlerFunc *http.HandlerFunc
}

// NewServer returns a mock server for unit testing purposes.
func NewServer() *ghttp.Server {
	var server *ghttp.Server
	server = ghttp.NewServer()

	server.RouteToHandler("POST", "/GetToken", func(w http.ResponseWriter,
		r *http.Request) {
		req, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()

		var login TokenLogin
		_ = json.Unmarshal(req, &login)
		if login.Username != validusername || login.Password != validpassword {
			w.Write([]byte(`error: Username and Password not valid for any
				licenses on this server`))
			return
		}

		w.Write([]byte(token))
		return
	})

	server.RouteToHandler("POST", "/MagicJson", func(w http.ResponseWriter,
		r *http.Request) {
		var magicbody MagicJsonBody

		reqbody, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()

		err := json.Unmarshal(reqbody, &magicbody)
		if err != nil {
			w.Write([]byte(`[{"Error":"Object reference not set to an instance
				of an object."}]`))
			return
		}

		if magicbody.Token != token {
			w.Write([]byte(`[{"Error":"Magic Error - As a Security precaution,
				you have been logged out due to inactivity."}]`))
			return
		}

		handler, err := getMockHandler(magicbody.Action)
		if err != nil {
			fmt.Println("no handler!")
			w.Write([]byte(`[{"Error":"Error: Action is not valid for this license."}]`))
			return
		}
		handler(w, r)
	})

	return server
}

func init() {
	validusername = os.Getenv("ALLSCRIPTS_USERNAME")
	validpassword = os.Getenv("ALLSCRIPTS_PASSWORD")
}
