package allscripts

import (
	"encoding/json"
	"os"
	"testing"
)

func TestGetScheduleBadIncludePic(t *testing.T) {
	server = testServer()
	client := NewClient("GRT", "TestApp", os.Getenv("ALLSCRIPTS_USERNAME"),
		os.Getenv("ALLSCRIPTS_PASSWORD"), server.URL())
	_, err := client.GetSchedule("", "", "nonsense", "", "")
	if err == nil {
		t.Errorf("Expected error thrown for invalid includepic passed")
	}
}

func TestGetScheduleBadSinceDate(t *testing.T) {
	server = testServer()
	client := NewClient("GRT", "TestApp", os.Getenv("ALLSCRIPTS_USERNAME"),
		os.Getenv("ALLSCRIPTS_PASSWORD"), server.URL())
	_, err := client.GetSchedule("", "nonsense", "Y", "", "")
	if err == nil {
		t.Errorf("Expected error thrown for invalid since date passed")
	}
}

func TestGetScheduleBadDateRange(t *testing.T) {
	server = testServer()
	client := NewClient("GRT", "TestApp", os.Getenv("ALLSCRIPTS_USERNAME"),
		os.Getenv("ALLSCRIPTS_PASSWORD"), server.URL())
	_, err := client.GetSchedule("nonsense", "", "Y", "", "")
	if err == nil {
		t.Errorf("Expected error thrown for invalid date range passed")
	}
}

func TestGetScheduleSuccess(t *testing.T) {
	server = testServer()
	client := NewClient("GRT", "TestApp", os.Getenv("ALLSCRIPTS_USERNAME"),
		os.Getenv("ALLSCRIPTS_PASSWORD"), server.URL())
	_, err := client.GetSchedule("7/1/2015|7/13/2015", "7/12/2015", "Y", "", "Ordered")
	if err != nil {
		t.Errorf("Expected no error for valid GetSchedule request")
	}
}

func TestGetScheduleNoDateRangeFailure(t *testing.T) {
	server = testServer()
	client := NewClient("GRT", "TestApp", os.Getenv("ALLSCRIPTS_USERNAME"),
		os.Getenv("ALLSCRIPTS_PASSWORD"), server.URL())
	_, err := client.GetSchedule("", "", "", "", "")
	if err == nil {
		t.Errorf("Expected error thrown for empty date range passed")
	}
}

func TestGetScheduleSuccessAllDefaults(t *testing.T) {
	server = testServer()
	client := NewClient("GRT", "TestApp", os.Getenv("ALLSCRIPTS_USERNAME"),
		os.Getenv("ALLSCRIPTS_PASSWORD"), server.URL())
	_, err := client.GetSchedule("7/1/2015|7/13/2015", "", "", "", "")
	if err != nil {
		t.Errorf("Expected no error for valid GetSchedule request")
	}
}

func TestGetScheduleSuccessFindingsReturned(t *testing.T) {
	server = testServer()
	client := NewClient("GRT", "TestApp", os.Getenv("ALLSCRIPTS_USERNAME"),
		os.Getenv("ALLSCRIPTS_PASSWORD"), server.URL())
	data, err := client.GetSchedule("7/1/2015|7/13/2015", "7/12/2015", "Y", "", "Ordered")
	if err != nil {
		t.Errorf("Expected no error for valid GetSchedule request: %s",
			err)
	}
	var appts Appointments
	err = json.Unmarshal(data, &appts)
	if err != nil {
		t.Errorf("Expected no error in unmarshaling valid GetSchedule request")
	}

	if len(appts) < 1 {
		t.Errorf("Expect that there will be at least one appointment")
	}
}
