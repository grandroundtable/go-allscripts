# go-allscripts

![TravisCI badge](https://api.travis-ci.org/grandroundtable/go-allscripts.svg)
[![GoDoc](https://godoc.org/github.com/grandroundtable/go-allscripts?status.svg)](http://godoc.org/github.com/grandroundtable/go-allscripts)

A Go library for accessing the Allscripts Professional Unity API. Visit 
[developers.allscripts.com](http://developer.allscripts.com) for more info on
the API.  

This library does not cover all calls to the Unity API, rather just the ones that we need at this time. We will accept pull requests for additional calls that you would like to see included. Please make sure to write tests for the functionality that you add.  

### Usage

```sh
client := allscripts.NewClient("appuserid", "appname", "appusername",
    "apppassword", "url")
raw, _ := client.GetClinicalSummary("patientid", "section")
var summary []allscripts.ClinicalSummary
_ = json.Unmarshal(raw, &summary)
```


### Tests
A mock server with mock JSON responses is included in the `mock` directory. Set environment variables `ALLSCRIPTS_USERNAME` and `ALLSCRIPTS_PASSWORD` and then run the tests with:  

`go test`

### Notes
* This package does not do any type casting. All Allscripts requests take string values and return string values.
* All sample data included with the mock server is from FAKE patients in the Allscripts sandbox.