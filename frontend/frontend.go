package frontend

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jcscottiii/testgaeapp/config"

	"appengine"
	"appengine/urlfetch"
)

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, err := client.Get(config.BaseURL + "/api")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "HTTP GET from API call returned: %s", string(response))
}

func init() {
	http.HandleFunc("/", handler)
}
