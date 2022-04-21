package main

import (
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8081", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	
	fmt.Fprintf(w, "Hello %s, I am %s years old!", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile("myfamily/family.txt")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	fmt.Fprintf(w, "My Family: %s.", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request) {

	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	
	fmt.Fprintf(w, "User: %s. Password: %s.", user, password)
}

func Healthz(w http.ResponseWriter, r *http.Request) {

	duration := time.Since(startedAt)

	if duration.Seconds() > 25 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Server is not ready. It has been up for %s.", duration)))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Server is ready. It has been up for %s.", duration)))
	}

}