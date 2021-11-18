package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/lekan-pvp/incr1/internal/app/shorter"
	"io"
	"log"
	"net/http"
	"strconv"
)

type URLs struct {
	Long string
	Short string
}

var shorts = make(map[int]URLs)

var ID = 1

func CreateShortURLHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id := &ID
	switch r.Method {
	case "POST":
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		long := string(body)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(201)
		short := shorter.Shorting(*id)
		shorts[*id] = URLs{
			Long: long,
			Short: short,
		}
		*id++
		w.Write([]byte(short))
	}
}

func GetURLById(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	param := r.URL.Path[len("http://localhost:8080/"):]
	id, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, "Wrong", 400)
		return
	}

	long := shorts[id].Long
	if long == "" {
		http.Error(w, "Wrong id", 400)
		return
	}
	w.WriteHeader(307)
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Location", long)
	w.Write([]byte(""))
}

func main() {
	router := httprouter.New()
	router.POST("/", CreateShortURLHandler)
	router.GET("/{id}", GetURLById)

	log.Fatal(http.ListenAndServe(":8080", router))
}
