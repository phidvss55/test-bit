package api

import (
	"context"
	"encoding/json"
	"fmt"
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
	http.HandleFunc("/healthz", s.handleHealthCheck)
	http.HandleFunc("/test", s.handleTest)
	http.HandleFunc("/", s.handleGetCatFact)
	
	// Product routes
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			s.handleGetAllProducts(w, r)
		} else if r.Method == http.MethodPost {
			s.handleCreateProduct(w, r)
		} else {
			writeJson(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		}
	})
	
	http.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			s.handleGetProduct(w, r)
		} else if r.Method == http.MethodPut {
			s.handleUpdateProduct(w, r)
		} else if r.Method == http.MethodDelete {
			s.handleDeleteProduct(w, r)
		} else {
			writeJson(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		}
	})

	fmt.Printf("API server listening on %s\n", listenAddress)

	return http.ListenAndServe(listenAddress, nil)
}

func (s *ApiServer) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	writeJson(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *ApiServer) handleTest(w http.ResponseWriter, r *http.Request) {
	writeJson(w, http.StatusOK, map[string]string{"status": "test is ok", "error": ""})
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
