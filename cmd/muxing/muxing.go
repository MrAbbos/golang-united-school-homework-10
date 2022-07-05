package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	r := mux.NewRouter()

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	r.HandleFunc("/", InitialHandler)
	r.HandleFunc("/name/{PARAM}", PathParamHandler)
	r.HandleFunc("/bad", InternalErrHandler)
	r.HandleFunc("/data", BodyReadHandler)
	r.HandleFunc("/headers", HeaderReadHandler)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), r); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func InitialHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is up and running")
}

func PathParamHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	parameter := params["PARAM"]
	// w.WriteHeader(http.StatusOK)
	bytes, _ := json.Marshal("Hello, " + parameter)
	fmt.Println(bytes)
	w.Write(bytes)

	// fmt.Fprintf(w, "Hello, %v\n", parameter)
}

func InternalErrHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

type body struct {
	PARAM string `json:"PARAM"`
}

func BodyReadHandler(w http.ResponseWriter, r *http.Request) {
	b := body{}
	json.NewDecoder(r.Body).Decode(&b)
	bytes, _ := json.Marshal("I got message:\n" + b.PARAM)
	w.Write(bytes)
}

func HeaderReadHandler(w http.ResponseWriter, r *http.Request) {
	a := r.Header.Get("a")
	b := r.Header.Get("b")
	c:=a+b
	bytes, _ := json.Marshal(`a+b: "`+c+`"`)
	w.Write(bytes)
}