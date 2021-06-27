package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"./internal"
	"github.com/gorilla/mux"
)

func main(){
	handleRequest()
}

func handleRequest() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/get/currency/all", getAllCurrecny).Methods("GET")
	router.HandleFunc("/get/currency/{currency}", getCurrecny).Methods("GET")
	router.HandleFunc("/", test).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func test(w http.ResponseWriter, r *http.Request) {
	
	enableCors(&w)
	response :="Test is working fine"
        fmt.Println(response)
		json.NewEncoder(w).Encode(response)	
			
}
func getAllCurrecny(w http.ResponseWriter, r *http.Request) {
	
	enableCors(&w)
	response :=internal.AllCurrencyDetail()
        fmt.Println(response)
		json.NewEncoder(w).Encode(response)	
			
}

func getCurrecny(w http.ResponseWriter, r *http.Request) {
	
	enableCors(&w)
	params := mux.Vars(r) // Get params
	currencyID :=params["currency"]
	response :=internal.CurrencyDetail(currencyID)
        fmt.Println(response)
		json.NewEncoder(w).Encode(response)	
			
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}