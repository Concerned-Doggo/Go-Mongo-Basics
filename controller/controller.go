package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Concerned-Doggo/mongoApi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb+srv://<username>:<password>@cluster0.eufsz.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const collectionName = "watchlist"

var collection *mongo.Collection

func init(){
    // client option
    clientOption := options.Client().ApplyURI(connectionString)

    // connect to mongo
    client, err := mongo.Connect(clientOption)
    if err != nil {
        log.Fatal("Error in connect to MONGODB: ", err)
    }
    fmt.Println("MongoDB connection success")

    collection = client.Database(dbName).Collection(collectionName)
    fmt.Println("collection reference/instance is ready")
}


func GetAllMovies(w http.ResponseWriter, r *http.Request){
    movies := getAllMovies()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(movies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
    movie := model.Netflix{}
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Allow-Control-Allow-Methods", "POST")

    json.NewDecoder(r.Body).Decode(&movie)
    insertOneMovie(movie)
    json.NewEncoder(w).Encode(movie)
}


func MarkAsWatched(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Allow-Control-Allow-Methods", "PUT")

    params := mux.Vars(r)
    movieId := params["id"]
    updateOneMovie(movieId)
    json.NewEncoder(w).Encode(movieId)
}


func DeleteOneMovie(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

    params := mux.Vars(r)
    movieId := params["id"]

    deleteOneMovie(movieId)
    json.NewEncoder(w).Encode(movieId)
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

    deleteAllMovies()
    message := "deleted all movies"
    json.NewEncoder(w).Encode(message)
}
