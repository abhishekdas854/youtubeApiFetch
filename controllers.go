package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func WriteJSON(w http.ResponseWriter, status int, val any) error {

	w.Header().Set("Contetnt-Type", "Application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(val)
}

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewApiServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	log.Println("Youtube data fetch api running on port: ", s.listenAddr)

	router.HandleFunc("/videosPaginated/{pageNo}", makeHttpHandleFunc(s.GetVideoDataPaginated))

	router.HandleFunc("/getVideo/{description}/{title}", makeHttpHandleFunc(s.GetVideoDataUsingTitleAndDescription))

	http.ListenAndServe(s.listenAddr, router)

}

func (s *APIServer) GetVideoDataPaginated(w http.ResponseWriter, r *http.Request) error {

	if r.Method != "GET" {
		return errors.New("Given endpoint only accepts get request")
	}

	return nil
}

func (s *APIServer) GetVideoDataUsingTitleAndDescription(w http.ResponseWriter, r *http.Request) error {
	return nil
}
