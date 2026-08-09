package api

import "net/http"

func (app *Application) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK - Health check passed"))
}
