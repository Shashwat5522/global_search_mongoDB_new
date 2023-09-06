package search

import (
	"encoding/json"
	"golang_globalsearch_new/apis/search/database"
	"log"
	"net/http"
)

type Handler struct {
	searchDB database.SearchDB
}

func NewHandler(searchDB database.SearchDB) Handler {
	return Handler{
		searchDB: searchDB,
	}
}

func (h *Handler) CreateObjects(w http.ResponseWriter, r *http.Request) {

	err := h.searchDB.CreateObjects()
	if err != nil {
		log.Fatalf("error is:=%v\n", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("objects created successfully!!!"))
}

func (h *Handler) SearchObjects(w http.ResponseWriter, r *http.Request) {
	searchWord := r.URL.Query()
	resp, err := h.searchDB.SearchObject(searchWord.Get("search"))
	if err != nil {
		log.Fatal(err)
	}
	response, jsonErr := json.Marshal(resp)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
