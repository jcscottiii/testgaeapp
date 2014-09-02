// +build appengine

package frontend
import (
	"net/http"

	"appengine"
	"appengine/urlfetch"
)

func SimpleOutsideCall(r *http.Request, w *http.ResponseWriter, baseurl string) (*http.Response, error) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, err := client.Get(baseurl + "/api")
	if err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	return resp, nil
}
