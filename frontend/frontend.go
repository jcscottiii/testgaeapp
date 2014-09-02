package frontend

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jcscottiii/testgaeapp/config"
)

func FrontendInit() {
	// This is just a placeholder function so that the standalone can properly
	// include this package. By being able to include this package, the init()
	// function will be called in the beginning like with using appengine and
	// register the handler(s).
}

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := SimpleOutsideCall(r, &w, config.BaseURL)
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
