package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var collection *mongo.Collection

func init(){
loadTheEnv()
createDBInstance()

}

func loadTheEnv(){

	err := godotenv.Load(".env")
	
	if err != nil{
		log.Fatal("Error loading the .env file")
	}
}

func createDBInstance(){
	connectionstring := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_COLLECTION_NAME")
	clientOptions := options.Client().ApplyURI(connectionstring)
	client, err := mongo.Connect(context.TODO(),clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(),nil)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb")

	client.Database(dbName).Collection(collName)
	fmt.Println("collection instance created")
}




func GetAllTasks(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Context-Type","application/x-www-form-urlencoded")

	w.Header().Set("Access-control-Allow-Origin","*")
	payload := getAllTasks()
	json.NewEncoder(w).Encode(payload)

}

func CreatTask(w http.ResponseWriter,r *http.Request){

}

func TaskComplete(w http.ResponseWriter,r *http.Request){

}
 
func UndoTask(w http.ResponseWriter,r *http.Request){



}

func DeleteTask(w http.ResponseWriter,r *http.Request){

}

func DeleteAllTasks(w http.ResponseWriter,r *http.Request){


}