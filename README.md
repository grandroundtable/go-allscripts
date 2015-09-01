# go-allscripts

A Go library for accessing the Allscripts Professional Unity API. Visit 
[developers.allscripts.com](http://developer.allscripts.com) for more info on
the API.  

### Usage

```sh
client := allscripts.NewClient("appuserid", "appname", "appusername",
    "apppassword", "url")
raw, _ := client.GetClinicalSummary("patientid", "section")
var summary []allscripts.ClinicalSummary
_ = json.Unmarshal(raw, &summary)
```

### Notes
* This package does not do any type casting. All Allscripts requests take string values and return string values.
* All sample data included with the mock server is from FAKE patients in the Allscripts sandbox.