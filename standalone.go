// +build !appengine

package main

import (
	"net/http"
	"github.com/jcscottiii/testgaeapp/backend"
	"github.com/jcscottiii/testgaeapp/frontend"
)

func main() {
	backend.BackendInit()
	frontend.FrontendInit()
	http.ListenAndServe("localhost:8080", nil)
}
