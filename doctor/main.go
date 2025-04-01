package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main(){
	r:= mux.NewRouter()
	r.HandleFunc("/doctors", GetDoctors).Methods("GET")
	r.HandleFunc("/doctor/{id}", GetDoctorByID).Methods("GET")
	r.HandleFunc("/doctor", CreateDoctor).Methods("POST")

	fmt.Println("Server Running on Port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
