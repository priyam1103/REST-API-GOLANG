package helper


import (
"github.com/gorilla/mux"
"context"
"encoding/json"
"time"
"net/http"
"github.com/priyam1103/REST-API/database"
"go.mongodb.org/mongo-driver/bson/primitive"
"go.mongodb.org/mongo-driver/bson"
)
type Person struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}
func CreatePerson(res http.ResponseWriter, req *http.Request){
	res.Header().Add("content-type","application/json")
	var person Person
	json.NewDecoder(req.Body).Decode(&person)
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.Klient.Database("godevp").Collection("people")
	result, err := collection.InsertOne(ctx,person)
	if err!= nil{
		
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message":"`+err.Error()+`"}`))
		return
	}	
	json.NewEncoder(res).Encode(result) 
	
}

func GetPerson (res http.ResponseWriter, req *http.Request){
	res.Header().Add("content-type","application/json")
	var people []Person
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.Klient.Database("godevp").Collection("people")
	cursor, err := collection.Find(ctx,bson.M{})
	if err!= nil{
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message":"`+err.Error()+`"}`))
		return
	}	
	defer cursor.Close(ctx)
	for cursor.Next(ctx){
		 var person Person
		 cursor.Decode(&person)
		 people=append(people,person)
	}
	json.NewEncoder(res).Encode(people)
}

func GetPersonExact (res http.ResponseWriter, req *http.Request){
	res.Header().Add("content-type","application/json")
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person Person
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.Klient.Database("godevp").Collection("people")
	err := collection.FindOne(ctx, Person{ID:id}).Decode(&person)
	if err!= nil{
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message":"`+err.Error()+`"}`))
		return
	}	
	json.NewEncoder(res).Encode(person)

}
func DeletePerson (res http.ResponseWriter , req *http.Request){
	res.Header().Add("content-type","application/json")
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.Klient.Database("godevp").Collection("people")
	resl, err := collection.DeleteOne(ctx,bson.M{"_id":id})
	if err!= nil{
		
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message":"`+err.Error()+`"}`))
		return
	}	
	json.NewEncoder(res).Encode(resl)
}

func UpdatePerson (res http.ResponseWriter, req *http.Request){
	res.Header().Add("content-type","application/json")
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.Klient.Database("godevp").Collection("people")
	var perso Person
	json.NewDecoder(req.Body).Decode(&perso)
	resl,err := collection.UpdateOne(ctx,bson.M{"_id":id},bson.D{{"$set",perso}})
	if err!= nil{
		
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message":"`+err.Error()+`"}`))
		return
	}	
	json.NewEncoder(res).Encode(resl)

}