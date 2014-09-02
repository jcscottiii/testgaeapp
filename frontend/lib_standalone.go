// +build !appengine

package frontend
import (
	"net/http"
)
func SimpleOutsideCall(r *http.Request, w *http.ResponseWriter, baseurl string) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Get(baseurl + "/api")
	if err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	return resp, nil
}
