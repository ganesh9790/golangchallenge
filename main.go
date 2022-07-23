package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type GetCurrencyCryptoPrices struct {
	Id          string `json:"id"`
	FullName    string `json:"fullName"`
	Ask         string `json:"ask"`
	Bid         string `json:"bid"`
	Last        string `json:"last"`
	Open        string `json:"open"`
	Low         string `json:"low"`
	High        string `json:"high"`
	FeeCurrency string `json:"feeCurrency"`
}

func main() {

	loadRoutes()

}

func loadRoutes() {
	r := mux.NewRouter()
	registerServiceRoutes(r)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln(err)
	}

}

func registerServiceRoutes(r *mux.Router) {

	setCurrencyRoutes(r)

}
func setCurrencyRoutes(r *mux.Router) {
	r.HandleFunc("/currency/BTCUSD", handleGetBtc).Methods(http.MethodGet)
	r.HandleFunc("/currency/ETHBTC", handleGetEth).Methods(http.MethodGet)
	r.HandleFunc("/currency/all", handleGetCurrency).Methods(http.MethodGet)
}

// handleGetBtc returns all the BTC crypto details.
func handleGetBtc(rw http.ResponseWriter, r *http.Request) {

	result := GetCurrencyCryptoPrices{}
	encodeResponse(rw, result)
	log.Println("Completed Processing for Eth")
}

// handleGetEth returns all the ETH crypto details.
func handleGetEth(rw http.ResponseWriter, r *http.Request) {
	result := GetCurrencyCryptoPrices{}
	encodeResponse(rw, result)
	log.Println("Completed Processing for Eth")
}

// handleGetCurrency returns all the crypto Currency.
func handleGetCurrency(rw http.ResponseWriter, r *http.Request) {

	result := []GetCurrencyCryptoPrices{}
	encodeResponse(rw, result)
	log.Println("Completed Processing to Receive all Currency")
}

// encodeResponse turns the response into json and writes it back to the ResponseWriter
func encodeResponse(w http.ResponseWriter, i interface{}) {
	err := json.NewEncoder(w).Encode(&i)
	if err != nil {
		fmt.Println(err)
	}
}
