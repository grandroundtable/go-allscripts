package mock

import (
	"fmt"
	"net/http"
)

// SendSchedule responses to the "GetSchedule" Allscripts API action with sample
// data.
func SendSchedule(sample []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(sample)
	}
}

func init() {
	csfile := "getschedule.json"
	gssample, err := LoadSampleJson(csfile)
	if err != nil {
		panic(fmt.Sprintf("%s could not be loaded...", csfile))
	}
	AddMockHandler(SendSchedule(gssample), "GetSchedule")
}
