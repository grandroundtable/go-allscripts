package allscripts

import (
	"encoding/json"
	"os"
	"testing"
)

func TestGetPatientSuccess(t *testing.T) {
	server = testServer()
	client := NewClient("GRT", "TestApp", os.Getenv("ALLSCRIPTS_USERNAME"),
		os.Getenv("ALLSCRIPTS_PASSWORD"), server.URL())
	_, err := client.GetPatient("212")
	if err != nil {
		t.Errorf("Expected no error for valid GetPatient request")
	}
}

func TestGetPatientSuccessFindingsReturned(t *testing.T) {
	server = testServer()
	client := NewClient("GRT", "TestApp", os.Getenv("ALLSCRIPTS_USERNAME"),
		os.Getenv("ALLSCRIPTS_PASSWORD"), server.URL())
	data, err := client.GetPatient("212")
	if err != nil {
		t.Errorf("Expected no error for valid GetPatient request: %s",
			err)
	}
	var patient Patient
	err = json.Unmarshal(data, &patient)
	if err != nil {
		t.Errorf("Expected no error in unmarshaling valid GetPatient request")
	}

	if patient.Lastname == "" {
		t.Errorf("Expect that a last name for the patient")
	}
}
