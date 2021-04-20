package main

import (
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", invalidHeaders)
	http.ListenAndServe(":8090", nil)
}

func invalidHeaders(w http.ResponseWriter, r *http.Request) {

	if len(os.Args) > 2 {
		w.Header().Set(os.Args[1], os.Args[2])
	} else {
		w.Header().Set("INVALID HEADAH", "invalid")
	}
}
