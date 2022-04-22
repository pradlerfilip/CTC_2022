package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Product struct {
	Name   string `json:"Name"`
	Value  int32  `json:"Value"`
	Amount int32  `json:"Amount"`
}

var Products []Product

func returnProducts(w http.ResponseWriter, r *http.Request) {
	//json.NewEncoder(w).Encode(Products)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.ykdp5.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("cv03")
	collection := database.Collection("products")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var products []bson.M
	if err = cursor.All(ctx, &products); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(products)
	fmt.Println(products)
}

func returnProductByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	/*
		for _, product := range Products {
			if product.Id == key {
				json.NewEncoder(w).Encode(product)
			}
		}
	*/
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.ykdp5.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("cv03")
	collection := database.Collection("products")
	var result bson.M
	err = collection.FindOne(ctx, bson.D{{"name", key}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			fmt.Println("No documents")
			return
		}
		panic(err)
	}
	fmt.Println(result)
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var product Product
	json.Unmarshal(reqBody, &product)
	//fmt.Println(product.Value)
	/*
		var product Product
		json.Unmarshal(reqBody, &product)

		Products = append(Products, product)

		json.NewEncoder(w).Encode(product)
	*/
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.ykdp5.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("cv03")
	collection := database.Collection("products")
	doc := bson.D{{"name", product.Name}, {"value", product.Value}, {"amount", product.Amount}}
	result, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	/*
		vars := mux.Vars(r)
		id := vars["id"]
		for index, product := range Products {
			if product.Id == id {
				Products = append(Products[:index], Products[index+1:]...)
			}
		}
	*/
	vars := mux.Vars(r)
	key := vars["id"]
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.ykdp5.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(key)
	database := client.Database("cv03")
	collection := database.Collection("products")
	result, err := collection.DeleteOne(context.TODO(), bson.M{"name": key})
	if err != nil {
		panic(err)
	}
	fmt.Println(result.DeletedCount)
}

func runClient(vw *sync.WaitGroup) {
	for true {
		//print("cau")
	}
	vw.Done()
}

func runServer(router *mux.Router, vw *sync.WaitGroup) {
	log.Fatal(http.ListenAndServe(":8080", router))
	vw.Done()
}

func main() {
	var vw sync.WaitGroup
	/*
		Products = []Product{
			{Id: "10000", Name: "Nvidia GTX 1060", Price: 8999, Amount: 10},
			{Id: "10001", Name: "Nvidia GTX 1070", Price: 12999, Amount: 15},
		}
	*/
	//product1 := bson.D{{"name", "NVIDIA GTX 1070"}, {"value", 6900}, {"amount", 10}}
	//product2 := bson.D{{"name", "NVIDIA GTX 1080"}, {"value", 12900}, {"amount", 5}}
	//_, err = collection.InsertOne(ctx, product1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//_, err = collection.InsertOne(ctx, product2)
	//if err != nil {
	//	log.Fatal(err)
	//}
	var router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/products", returnProducts)
	router.HandleFunc("/product/{id}", returnProductByName)
	router.HandleFunc("/product", addProduct).Methods("POST")
	router.HandleFunc("/product/{id}", deleteProduct)

	vw.Add(1)
	go runServer(router, &vw)
	vw.Add(1)
	go runClient(&vw)
	vw.Wait()
}
