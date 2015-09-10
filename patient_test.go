package allscripts

import (
	"encoding/json"
	"os"
	"sync"
	"testing"

	"github.com/grandroundtable/go-allscripts/mock"
	"github.com/onsi/gomega/ghttp"
)

var (
	server *ghttp.Server
	sOnce  sync.Once
)

func testServer() *ghttp.Server {
	sOnce.Do(func() {
		server = mock.NewServer()
	})

	return server
}

func TestClinicalSummaryBadLogin(t *testing.T) {
	server = testServer()
	client := NewClient("GRT", "GRTTEST", "marlowe", "gunz", server.URL())
	_, err := client.GetClinicalSummary("213", "list")
	if err == nil {
		t.Errorf("Expected error thrown for incorrect login credentials")
	}
}

func TestClinicalSummaryBadSection(t *testing.T) {
	server = testServer()
	client := NewClient("GRT", "TestApp", os.Getenv("ALLSCRIPTS_USERNAME"),
		os.Getenv("ALLSCRIPTS_PASSWORD"), server.URL())
	_, err := client.GetClinicalSummary("212", "nonsense")
	if err == nil {
		t.Errorf("Expected error thrown for invalid section passed")
	}
}

func TestClinicalSummarySuccess(t *testing.T) {
	server = testServer()
	client := NewClient("GRT", "TestApp", os.Getenv("ALLSCRIPTS_USERNAME"),
		os.Getenv("ALLSCRIPTS_PASSWORD"), server.URL())
	_, err := client.GetClinicalSummary("212", "list")
	if err != nil {
		t.Errorf("Expected no error for valid GetClinicalSummary request")
	}
}

func TestClinicalSummarySuccessFindingsReturned(t *testing.T) {
	server = testServer()
	client := NewClient("GRT", "TestApp", os.Getenv("ALLSCRIPTS_USERNAME"),
		os.Getenv("ALLSCRIPTS_PASSWORD"), server.URL())
	data, err := client.GetClinicalSummary("212", "list")
	if err != nil {
		t.Errorf("Expected no error for valid GetClinicalSummary request: %s",
			err)
	}
	var summary []ClinicalSummary
	err = json.Unmarshal(data, &summary)
	if err != nil {
		t.Errorf("Expected no error in unmarshaling valid GetClinicalSummary request")
	}

	if len(summary[0].Findings) < 1 {
		t.Errorf("Expect that there will be at least one clinical finding")
	}
}
