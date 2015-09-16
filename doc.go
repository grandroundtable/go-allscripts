/*
A Go library for accessing the Allscripts Professional Unity API.

Example:

    package main

    import (
        "github.com/grandroundtable/go-allscripts"
    )

    func main() {
        client := allscripts.NewClient("appuserid", "appname", "appusername",
            "apppassword", "url")
        raw, _ := client.GetClinicalSummary("patientid", "section")
        var summary []allscripts.ClinicalSummary
        _ = json.Unmarshal(raw, &summary)
    }
*/
package allscripts
