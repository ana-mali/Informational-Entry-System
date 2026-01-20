package notes

import "net/http"

func Router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /notes", list)
	mux.HandleFunc("POST /notes", add)
	mux.HandleFunc("DELETE /notes/{id}", delete)

	return mux
}
