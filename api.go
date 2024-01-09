package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func WriteJson(w http.ResponseWriter, status int, value any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(value)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error  string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type ApiServer struct {
	ListenAddr  string
}


func NewApiServer(ListenAddr string) *ApiServer {
	return &ApiServer{
		ListenAddr: ListenAddr,
	}
}


func (s *ApiServer) Run() {
	router := mux.NewRouter()

	router.Handle("/account", makeHTTPHandleFunc(s.handleAccount))
	router.Handle("/account/{id}", makeHTTPHandleFunc(s.handleAccount))

	log.Printf("HTTP SERVER IS RUNNING ON PORT %s \n", s.ListenAddr)

	http.ListenAndServe(s.ListenAddr, router)

}

func (s *ApiServer) handleAccount(w http.ResponseWriter, r * http.Request) error {

	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	
	return fmt.Errorf("method not allowed %s ", r.Method)
}


func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r * http.Request) error {
	id := mux.Vars(r)["id"]
	println(id)
	return WriteJson(w, http.StatusOK, Account{})
}



func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, r * http.Request) error {
	return nil
}




func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r * http.Request) error {
	return nil
}

func (s *ApiServer) handleTransferAccount(w http.ResponseWriter, r * http.Request) error {
	return nil
}

