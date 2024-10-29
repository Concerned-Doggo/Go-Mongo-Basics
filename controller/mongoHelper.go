package controller

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Concerned-Doggo/mongoApi/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func insertOneMovie(movie model.Netflix){
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    res, err := collection.InsertOne(ctx, movie)
    if err != nil{
        log.Fatal("error in inserting one movie: ", err)
    }
    
    id := res.InsertedID
    fmt.Println("inserted one movie with id: ", id)

    return
}


func updateOneMovie(movieId string){
    id, err := bson.ObjectIDFromHex(movieId)
    if err != nil{
        log.Fatal("error in converting string to ObjectID: ", err)
    }

    filter := bson.M{"_id":id}
    update := bson.M{"$set": bson.M{"watched":true}}

    result, err := collection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        log.Fatal("error in updating one movie", err)
    }

    fmt.Println("modified count: ", result.ModifiedCount)
}


func deleteOneMovie(movieId string){
    id, err := bson.ObjectIDFromHex(movieId)
    if err != nil {
        log.Fatal("error in converting string to ObjectID (deleteOneMovie): ", err)
    }

    filter := bson.M{"_id": id}
    result, err := collection.DeleteOne(context.Background(), filter)
    if err != nil{
        log.Fatal("error in deleting one movie: ", err)
    }

    fmt.Println("deleted count: ", result.DeletedCount)
} 

func deleteAllMovies(){
    result, err := collection.DeleteMany(context.Background(), bson.D{{}})
    if err != nil{
        log.Fatal("error in deleting all movies: ", err)
    }
    
    fmt.Println("deleted count: ", result.DeletedCount)
}

func getAllMovies() []bson.M{
    // cursor is a object of all the movies
    cursor, err := collection.Find(context.Background(), bson.D{{}})
    if err != nil{
        log.Fatal("error in finding all movies: ", err)
    }

    defer cursor.Close(context.Background())

    var movies []bson.M
    for cursor.Next(context.Background()){
        var movie bson.M
        err := cursor.Decode(&movie)
        if err != nil{
            log.Fatal("error in finding all movies: ", err)
        }
        movies = append(movies, movie)
    }
    return movies
}
