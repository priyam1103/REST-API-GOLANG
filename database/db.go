package database
import (
	
"go.mongodb.org/mongo-driver/mongo"
"fmt"
"time"
"context"
"go.mongodb.org/mongo-driver/mongo/options"

)

var Klient *mongo.Client 

func Createe() {


 clientOptions :=(options.Client().ApplyURI("mongodb://localhost:27017/"))
 
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
   Klient , _ = mongo.Connect(ctx, clientOptions)
fmt.Println("Connected")
	
}
