package main

import (
"github.com/gorilla/mux"
"go.mongodb.org/mongo-driver/mongo"
"fmt"
"net/http"
"context"
"time"
"github.com/priyam1103/REST-API/helper"
"github.com/priyam1103/REST-API/database"
)

var client *mongo.Client

func main(){
fmt.Println("Connected")
ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
database.Createe() 
fmt.Println("Connected to MongoDB!")
router := mux.NewRouter()
router.HandleFunc("/person",helper.CreatePerson).Methods("POST")
router.HandleFunc("/getperson",helper.GetPerson).Methods("GET")
router.HandleFunc("/getperson/{id}",helper.GetPersonExact).Methods("GET")
router.HandleFunc("/deleteperson/{id}",helper.DeletePerson).Methods("DELETE")
router.HandleFunc("/updateperson/{id}",helper.UpdatePerson).Methods("PUT")
http.ListenAndServe(":3000",router)
defer database.Klient.Disconnect(ctx)
}