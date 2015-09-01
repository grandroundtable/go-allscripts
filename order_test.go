package allscripts

import (
	"encoding/json"
	"os"
	"testing"
)

func TestGetOrdersBadStatus(t *testing.T) {
	server = testServer()
	client := NewClient("GRT", "TestApp", os.Getenv("ALLSCRIPTS_USERNAME"),
		os.Getenv("ALLSCRIPTS_PASSWORD"), server.URL())
	_, err := client.GetOrders("", "2015/08/16", "", "nonsense")
	if err == nil {
		t.Errorf("Expected error thrown for invalid status passed")
	}
}

func TestGetOrdersBadSinceDate(t *testing.T) {
	server = testServer()
	client := NewClient("GRT", "TestApp", os.Getenv("ALLSCRIPTS_USERNAME"),
		os.Getenv("ALLSCRIPTS_PASSWORD"), server.URL())
	_, err := client.GetOrders("", "2/08/16", "", "nonsense")
	if err == nil {
		t.Errorf("Expected error thrown for invalid since date passed")
	}
}

func TestGetOrdersSuccess(t *testing.T) {
	server = testServer()
	client := NewClient("GRT", "TestApp", os.Getenv("ALLSCRIPTS_USERNAME"),
		os.Getenv("ALLSCRIPTS_PASSWORD"), server.URL())
	_, err := client.GetOrders("", "2015/08/16", "", "Ordered")
	if err != nil {
		t.Errorf("Expected no error for valid GetOrders request")
	}
}

func TestGetOrdersSuccessAllDefaults(t *testing.T) {
	server = testServer()
	client := NewClient("GRT", "TestApp", os.Getenv("ALLSCRIPTS_USERNAME"),
		os.Getenv("ALLSCRIPTS_PASSWORD"), server.URL())
	_, err := client.GetOrders("", "", "", "")
	if err != nil {
		t.Errorf("Expected no error for valid GetOrders request")
	}
}

func TestGetOrdersSuccessFindingsReturned(t *testing.T) {
	server = testServer()
	client := NewClient("GRT", "TestApp", os.Getenv("ALLSCRIPTS_USERNAME"),
		os.Getenv("ALLSCRIPTS_PASSWORD"), server.URL())
	data, err := client.GetOrders("", "2015/08/16", "", "Ordered")
	if err != nil {
		t.Errorf("Expected no error for valid GetOrders request: %s",
			err)
	}
	var orders Orders
	err = json.Unmarshal(data, &orders)
	if err != nil {
		t.Errorf("Expected no error in unmarshaling valid GetOrders request")
	}

	if len(orders) < 1 {
		t.Errorf("Expect that there will be at least one order")
	}
}
