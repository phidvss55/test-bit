package api

import (
	"log"
)

type ApiServer struct {
	addr string
	db *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{addr: addr, db: db}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api").Subrouter()

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	log.Printf("Starting API server on %s", s.addr)
	return http.ListenAndServe(s.addr, router)
}