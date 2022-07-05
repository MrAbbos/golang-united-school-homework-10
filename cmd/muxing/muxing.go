package main

import (
	"fmt"
	"io"
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
	// bytes, _ := json.Marshal("Hello, " + parameter)
	fmt.Fprintf(w, "Hello, %v!", parameter)
	// fmt.Println(bytes)
	// w.Write(bytes)
}

func InternalErrHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

type body struct {
	PARAM string `json:"PARAM"`
}

func BodyReadHandler(w http.ResponseWriter, r *http.Request) {
	// b := body{}
	// json.NewDecoder(r.Body).Decode(&b)
	// bytes, _ := json.Marshal("I got message:\n" + b.PARAM)
	// w.Write(bytes)
	data, _ := io.ReadAll(r.Body)
	fmt.Fprintf(w, "I got message:\n%v", string(data))
}

func HeaderReadHandler(w http.ResponseWriter, r *http.Request) {
	a := r.Header.Get("a")
	b := r.Header.Get("b")
	inta, _ := strconv.Atoi(a)
	intb, _ := strconv.Atoi(b)
	c := inta + intb
	w.Header().Set("a+b", strconv.Itoa(c))
}
