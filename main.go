package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to GO API Testing out!")
}

func main(){
	r:=mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	fmt.Println("Server listening on Port:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}