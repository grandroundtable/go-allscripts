package mock

import (
	"fmt"
	"net/http"
)

// SendClinicalSummary responses to the "GetClinicalSummary" Allscripts API
// action with sample data.
func SendClinicalSummary(sample []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(sample)
	}
}

func init() {
	var (
		csfile = "getclinicalsummarysample.json"
	)
	gcssample, err := LoadSampleJson(csfile)
	if err != nil {
		panic(fmt.Sprintf("%s could not be loaded...", csfile))
	}
	AddMockHandler(SendClinicalSummary(gcssample), "GetClinicalSummary")
}
