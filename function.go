package Cloud_Functions_in_Go

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet, http.MethodPost:
		fmt.Fprintf(w, "Hello World: %s", r.Method)
	default:
		http.Error(w, "405 - Method not allowed", http.StatusMethodNotAllowed)
	}
}
