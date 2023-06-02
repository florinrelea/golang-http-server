package main

import (
	"log"
	"main/httpUtils"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiError struct {
	Error string
}

type APIFunction func(http.ResponseWriter, *http.Request) error

type APIServer struct {
	listenAddr string
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", adaptToHttpFunc((s.handleAccount)))

	log.Println("Server started on", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func adaptToHttpFunc(f APIFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			httpUtils.WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func CreateNewApiServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) handleAccount(writer http.ResponseWriter, reader *http.Request) error {
	if reader.Method == "GET" {
		return s.handleGetAccount(writer, reader)
	}

	return nil
}

func (s *APIServer) handleGetAccount(writer http.ResponseWriter, reader *http.Request) error {
	return nil
}
