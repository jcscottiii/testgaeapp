package backend

import (
	"fmt"
	"net/http"
)

func BackendInit() {
	// This is just a placeholder function so that the standalone can properly
	// include this package. By being able to include this package, the init()
	// function will be called in the beginning like with using appengine and
	// register the handler(s).
}

func init() {
	http.HandleFunc("/api", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World from James Scott!")
}
