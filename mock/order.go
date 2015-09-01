package mock

import (
	"fmt"
	"net/http"
)

// SendOrders responses to the "GetOrders" Allscripts API action with sample
// data.
func SendOrders(sample []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(sample)
	}
}

func init() {
	csfile := "getorders.json"
	gosample, err := LoadSampleJson(csfile)
	if err != nil {
		panic(fmt.Sprintf("%s could not be loaded...", csfile))
	}
	AddMockHandler(SendOrders(gosample), "GetOrders")
}
