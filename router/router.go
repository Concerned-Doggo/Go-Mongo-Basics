package router

import (

	"github.com/Concerned-Doggo/mongoApi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/getMovies", controller.GetAllMovies).Methods("GET")
    router.HandleFunc("/createMovie", controller.CreateMovie).Methods("POST")
    router.HandleFunc("/updateMovie/{id}", controller.MarkAsWatched).Methods("PUT")
    router.HandleFunc("/deleteMovies", controller.DeleteAllMovies).Methods("DELETE")
    router.HandleFunc("/deleteMovie/{id}", controller.DeleteOneMovie).Methods("DELETE")

    return router;
}
