package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Product struct {
	Name   string `bson:"Name" json:"Name"`
	Value  int32  `bson:"Value" json:"Value"`
	Amount int32  `bson:"Amount" json:"Amount"`
}

var collection *mongo.Collection

func GetProducts(response http.ResponseWriter, r *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var products []Product
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	}
	cursor.All(ctx, &products)
	fmt.Println(products)
	json.NewEncoder(response).Encode(products)
}

func GetProduct(response http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["Name"]
	response.Header().Set("Content-Type", "application/json")
	var product bson.M
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	collection.FindOne(ctx, bson.D{{"Name", key}}).Decode(&product)
	json.NewEncoder(response).Encode(product)
}

func AddProduct(response http.ResponseWriter, r *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var product Product
	json.Unmarshal(reqBody, &product)
	context.WithTimeout(context.Background(), 30*time.Second)
	doc := bson.D{{"Name", product.Name}, {"Value", product.Value}, {"Amount", product.Amount}}
	result, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func DeleteProduct(response http.ResponseWriter, r *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["Name"]
	context.WithTimeout(context.Background(), 30*time.Second)
	filter := bson.D{{"Name", key}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.DeletedCount)

}

func UpdateProduct(response http.ResponseWriter, r *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["Name"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var product Product
	json.Unmarshal(reqBody, &product)
	context.WithTimeout(context.Background(), 30*time.Second)
	filter := bson.D{{"Name", key}}
	update := bson.D{{"$set", bson.D{{"Name", product.Name}, {"Value", product.Value}, {"Amount", product.Amount}}}}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.MatchedCount)
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.ykdp5.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	collection = client.Database("cv03").Collection("products")
	router := mux.NewRouter()
	router.HandleFunc("/products", GetProducts).Methods("GET")
	router.HandleFunc("/product/{Name}", GetProduct).Methods("GET")
	router.HandleFunc("/product", AddProduct).Methods("POST")
	router.HandleFunc("/product/{Name}", DeleteProduct).Methods("DELETE")
	router.HandleFunc("/product/{Name}", UpdateProduct).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", router))
}
