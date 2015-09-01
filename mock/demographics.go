package mock

import (
	"fmt"
	"net/http"
)

// SendPatient responses to the "GetPatient" Allscripts API action with sample
// data.
func SendPatient(sample []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(sample)
	}
}

func init() {
	csfile := "getpatient.json"
	gpsample, err := LoadSampleJson(csfile)
	if err != nil {
		panic(fmt.Sprintf("%s could not be loaded...", csfile))
	}
	AddMockHandler(SendPatient(gpsample), "GetPatient")
}
