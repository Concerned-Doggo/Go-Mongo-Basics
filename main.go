package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Concerned-Doggo/mongoApi/router"
)

func main(){
    fmt.Println("server is starting...")
    router := router.Router()
    log.Fatal(http.ListenAndServe(":8080", router))
    fmt.Println("listing on port 8080...")
}
