package backend

import (
	"fmt"
	"net/http"
)

// ALL Paths end with "/". Be consistent...

// HomeHandler ...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "welcome home")
}

// HelloHandler ...
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "hello frontend")
}

// ApiTestHandler ...
func ApiTestHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/test/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "HI there from Go Web Svc. %s", r.URL.Path[1:])
	fmt.Printf("called Go Web Svc. %s", r.URL.Path[1:])
}

// errorHandler ...
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}
