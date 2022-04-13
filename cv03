package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type Product struct {
	Id     string  `json:"Id"`
	Name   string  `json:"Name"`
	Price  float64 `json:"Price"`
	Amount int32   `json:"Amount"`
}

var Products []Product

func returnProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Products)
}

func returnProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, product := range Products {
		if product.Id == key {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var product Product
	json.Unmarshal(reqBody, &product)

	Products = append(Products, product)

	json.NewEncoder(w).Encode(product)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for index, product := range Products {
		if product.Id == id {
			Products = append(Products[:index], Products[index+1:]...)
		}
	}
}

func runClient() {
	for true {
		print("cau")
	}
}

func runServer(router *mux.Router) {
	go log.Fatal(http.ListenAndServe(":8080", router))

}

func main() {
	var vw sync.WaitGroup
	Products = []Product{
		{Id: "10000", Name: "Nvidia GTX 1060", Price: 8999, Amount: 10},
		{Id: "10001", Name: "Nvidia GTX 1070", Price: 12999, Amount: 15},
	}
	var router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/products", returnProducts)
	router.HandleFunc("/product/{id}", returnProductByID)
	router.HandleFunc("/product", addProduct).Methods("POST")
	router.HandleFunc("/article/{id}", deleteProduct).Methods("DELETE")

	vw.Add(1)
	//go log.Fatal(http.ListenAndServe(":8080", router))
	go runServer(router)
	vw.Add(1)
	go runClient()
	vw.Wait()
}
