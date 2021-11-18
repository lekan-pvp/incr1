package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"strconv"
)

const prefix = "localhost:8080/"

type URLs struct {
	Long string
	Short string
}

var shorts = make(map[int]URLs)

var ID = 1

func CreateShortURLHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := &ID
	long, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(200)
	short := fmt.Sprintf("%s%d", prefix, id)
	shorts[*id] = URLs{
		Long: string(long),
		Short: short,
	}
	*id++
	_, err = w.Write([]byte(short))
}

func GetShortURLById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))

	log.Println(id)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(307)
	w.Header().Set("Location", shorts[id].Long)
	_, err = w.Write([]byte(""))
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func main() {
	router := httprouter.New()
	router.POST("/", CreateShortURLHandler)
	router.GET("/:id", GetShortURLById)

	log.Fatal(http.ListenAndServe(":8080", router))
}
