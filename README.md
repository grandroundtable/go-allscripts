# go-allscripts

A Go library for working with the Allscripts Professional Unity API. Visit 
[developers.allscripts.com](http://developer.allscripts.com) for more info on
the API.  

### Usage

```sh
client := allscripts.NewClient("appuserid", "appname", "appusername",
    "apppassword", "url")
raw, _ := client.GetClinicalSummary("patientid", "section")
var summary []allscripts.ClinicalSummary
_ = json.Unmarshal(raw, %summary)
```
