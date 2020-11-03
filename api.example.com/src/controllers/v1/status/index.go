package status

import "net/http"

// Index function
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("V1 Status is live!"))
}
