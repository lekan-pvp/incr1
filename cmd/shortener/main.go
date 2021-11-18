package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func CreateShortURLHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func GetShortURLById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func main() {
	router := httprouter.New()
	router.POST("/", CreateShortURLHandler)
	router.GET("/:id", GetShortURLById)

	log.Fatal(http.ListenAndServe(":8080", router))
}
