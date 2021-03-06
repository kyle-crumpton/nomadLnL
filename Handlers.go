package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func LetsAddNumbers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	numOne := vars["numone"]
	numTwo := vars["numtwo"]
	fmt.Fprintln(w, "Sum: ", numOne+numTwo)
}

func Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Version: 0.0.1")
}
