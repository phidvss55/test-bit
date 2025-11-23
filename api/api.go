package api

import (
	"context"
	"encoding/json"
	"go-circleci/services"
	"net/http"
)

type ApiServer struct {
	svc services.Service
}

func NewApiServer(svc services.Service) *ApiServer {
	return &ApiServer{svc: svc}
}

func (s *ApiServer) Start(listenAddress string) error {
	http.HandleFunc("/", s.handleGetCatFact)

	return http.ListenAndServe(listenAddress, nil)
}

func (s *ApiServer) handleGetCatFact(w http.ResponseWriter, r *http.Request) {
	fact, err := s.svc.GetCatFact(context.Background())
	if err != nil {
		writeJson(w, http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
		return
	}
	writeJson(w, http.StatusOK, fact)
}

func writeJson(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
