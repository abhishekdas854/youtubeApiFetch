package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

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

	router.HandleFunc("/getVideo", makeHttpHandleFunc(s.GetVideoDataUsingTitleAndDescription))

	http.ListenAndServe(s.listenAddr, router)

}

func (s *APIServer) GetVideoDataPaginated(w http.ResponseWriter, r *http.Request) error {

	if r.Method != "GET" {
		return errors.New("given endpoint only accepts get request")
	}

	pageNo := mux.Vars(r)["pageNo"]

	log.Println("Page no is: ", pageNo)
	log.Printf("Page no type is: %T\n", pageNo)

	pageNum, err := strconv.Atoi(pageNo)

	if err != nil {
		log.Fatal("Not valid page number: ", err)
	}

	res, err := s.store.GetDetailsFromDbPaginated(pageNum)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, res)
}

func (s *APIServer) GetVideoDataUsingTitleAndDescription(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return errors.New("given endpoint only accepts get request")
	}

	queryParams := r.URL.Query()

	description := queryParams.Get("description")
	title := queryParams.Get("title")

	res, err := s.store.GetDetailsUsingTitleAndDescription(title, description)

	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, res)
}
